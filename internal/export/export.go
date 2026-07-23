package export

import (
	"context"
	"fmt"
	"os"
	"strings"
	"time"

	"google.golang.org/api/drive/v3"
	"google.golang.org/api/option"
	"google.golang.org/api/sheets/v4"
)

const reportSheetTitle = "Flash P&L"

// NewDriveClient builds a Drive API client from the same service account key
// used for Sheets, scoped to files the account itself creates (drive.file) —
// enough to share the temp report spreadsheets this package creates.
func NewDriveClient(ctx context.Context) (*drive.Service, error) {
	raw := os.Getenv("GOOGLE_SERVICE_ACCOUNT_JSON")
	if raw == "" {
		return nil, fmt.Errorf("GOOGLE_SERVICE_ACCOUNT_JSON is not set")
	}
	return drive.NewService(ctx,
		option.WithCredentialsJSON([]byte(raw)),
		option.WithScopes(drive.DriveFileScope),
	)
}

func (b *builder) toRequests(sheetID int64) []*sheets.Request {
	var reqs []*sheets.Request
	for _, s := range b.styles {
		cf := &sheets.CellFormat{}
		var fields []string
		if s.bg != nil {
			cf.BackgroundColor = s.bg
			fields = append(fields, "userEnteredFormat.backgroundColor")
		}
		if s.fontColor != nil || s.bold || s.fontSize > 0 {
			tf := &sheets.TextFormat{Bold: s.bold, FontSize: s.fontSize}
			if s.fontColor != nil {
				tf.ForegroundColor = s.fontColor
			}
			cf.TextFormat = tf
			fields = append(fields, "userEnteredFormat.textFormat")
		}
		if s.numberFormat != "" {
			cf.NumberFormat = &sheets.NumberFormat{Type: "NUMBER", Pattern: s.numberFormat}
			fields = append(fields, "userEnteredFormat.numberFormat")
		}
		if len(fields) == 0 {
			continue
		}
		reqs = append(reqs, &sheets.Request{
			RepeatCell: &sheets.RepeatCellRequest{
				Range: &sheets.GridRange{
					SheetId: sheetID, StartRowIndex: int64(s.row), EndRowIndex: int64(s.row + 1),
					StartColumnIndex: int64(s.colStart), EndColumnIndex: int64(s.colEnd),
				},
				Cell:   &sheets.CellData{UserEnteredFormat: cf},
				Fields: strings.Join(fields, ","),
			},
		})
	}
	for _, m := range b.merges {
		reqs = append(reqs, &sheets.Request{
			MergeCells: &sheets.MergeCellsRequest{
				Range: &sheets.GridRange{
					SheetId: sheetID, StartRowIndex: int64(m.row), EndRowIndex: int64(m.row + 1),
					StartColumnIndex: int64(m.colStart), EndColumnIndex: int64(m.colEnd),
				},
				MergeType: "MERGE_ALL",
			},
		})
	}
	return reqs
}

func colWidthReq(sheetID int64, colIndex int, px int64) *sheets.Request {
	return &sheets.Request{
		UpdateDimensionProperties: &sheets.UpdateDimensionPropertiesRequest{
			Range:      &sheets.DimensionRange{SheetId: sheetID, Dimension: "COLUMNS", StartIndex: int64(colIndex), EndIndex: int64(colIndex + 1)},
			Properties: &sheets.DimensionProperties{PixelSize: px},
			Fields:     "pixelSize",
		},
	}
}

// Export ports exportFlashPL(): builds the formatted report into a brand new
// spreadsheet (owned by the service account) and shares it with userEmail so
// they can open the returned .xlsx export link.
func Export(ctx context.Context, sheetsSvc *sheets.Service, driveSvc *drive.Service, sheetTZ, origSheetName string, month, year int, pl PLData, userEmail string) (string, error) {
	loc, err := time.LoadLocation(sheetTZ)
	if err != nil {
		loc = time.UTC
	}
	generatedAt := time.Now().In(loc).Format("02 Jan 2006 15:04")

	b := buildReport(month, year, origSheetName, generatedAt, pl)
	title := fmt.Sprintf("Flash P&L - %s %d - %s", monthNames[month], year, origSheetName)

	newSS, err := sheetsSvc.Spreadsheets.Create(&sheets.Spreadsheet{
		Properties: &sheets.SpreadsheetProperties{Title: title},
		Sheets:     []*sheets.Sheet{{Properties: &sheets.SheetProperties{Title: reportSheetTitle}}},
	}).Context(ctx).Do()
	if err != nil {
		return "", fmt.Errorf("create report spreadsheet: %w", err)
	}
	ssID := newSS.SpreadsheetId
	sheetID := newSS.Sheets[0].Properties.SheetId

	if _, err := sheetsSvc.Spreadsheets.Values.Update(ssID, fmt.Sprintf("'%s'!A1", reportSheetTitle), &sheets.ValueRange{Values: b.values}).
		ValueInputOption("RAW").Context(ctx).Do(); err != nil {
		return "", fmt.Errorf("write report values: %w", err)
	}

	reqs := b.toRequests(sheetID)
	reqs = append(reqs,
		colWidthReq(sheetID, 0, 340),
		colWidthReq(sheetID, 1, 180),
		colWidthReq(sheetID, 2, 80),
	)
	if _, err := sheetsSvc.Spreadsheets.BatchUpdate(ssID, &sheets.BatchUpdateSpreadsheetRequest{Requests: reqs}).Context(ctx).Do(); err != nil {
		return "", fmt.Errorf("format report: %w", err)
	}

	if userEmail != "" {
		if _, err := driveSvc.Permissions.Create(ssID, &drive.Permission{
			Type: "user", Role: "reader", EmailAddress: userEmail,
		}).SendNotificationEmail(false).Context(ctx).Do(); err != nil {
			return "", fmt.Errorf("share report with %s: %w", userEmail, err)
		}
	}

	return fmt.Sprintf("https://docs.google.com/spreadsheets/d/%s/export?format=xlsx", ssID), nil
}
