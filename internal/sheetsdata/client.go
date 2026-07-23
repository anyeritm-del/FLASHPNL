package sheetsdata

import (
	"context"
	"fmt"
	"os"

	"google.golang.org/api/option"
	"google.golang.org/api/sheets/v4"
)

// NewClient builds a Sheets API client from the service account key in
// GOOGLE_SERVICE_ACCOUNT_JSON. Scope is read/write on spreadsheets only —
// Drive sharing (for exports) is handled by the export package with its own client.
func NewClient(ctx context.Context) (*sheets.Service, error) {
	raw := os.Getenv("GOOGLE_SERVICE_ACCOUNT_JSON")
	if raw == "" {
		return nil, fmt.Errorf("GOOGLE_SERVICE_ACCOUNT_JSON is not set")
	}
	return sheets.NewService(ctx,
		option.WithCredentialsJSON([]byte(raw)),
		option.WithScopes(sheets.SpreadsheetsScope),
	)
}

func SheetID() (string, error) {
	id := os.Getenv("GOOGLE_SHEET_ID")
	if id == "" {
		return "", fmt.Errorf("GOOGLE_SHEET_ID is not set")
	}
	return id, nil
}

// SpreadsheetName mirrors ss.getName() used for hotelName in the original initApp.
func SpreadsheetName(ctx context.Context, svc *sheets.Service, sheetID string) (string, error) {
	meta, err := svc.Spreadsheets.Get(sheetID).Fields("properties.title").Context(ctx).Do()
	if err != nil {
		return "", err
	}
	return meta.Properties.Title, nil
}

// TimeZone returns the spreadsheet's IANA timezone, used to format saved-at
// timestamps the same way Utilities.formatDate did under Apps Script.
func TimeZone(ctx context.Context, svc *sheets.Service, sheetID string) (string, error) {
	meta, err := svc.Spreadsheets.Get(sheetID).Fields("properties.timeZone").Context(ctx).Do()
	if err != nil {
		return "", err
	}
	return meta.Properties.TimeZone, nil
}
