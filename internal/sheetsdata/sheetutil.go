package sheetsdata

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"google.golang.org/api/googleapi"
	"google.golang.org/api/sheets/v4"
)

// isMissingSheetErr reports whether err is the Sheets API's "Unable to parse
// range" error, i.e. the referenced sheet tab doesn't exist yet.
func isMissingSheetErr(err error) bool {
	var apiErr *googleapi.Error
	if !errors.As(err, &apiErr) {
		return false
	}
	return apiErr.Code == 400 && strings.Contains(apiErr.Message, "Unable to parse range")
}

// colLetter converts a 1-indexed column number to its A1 letter(s), e.g. 1->A, 27->AA.
func colLetter(n int) string {
	var s string
	for n > 0 {
		n--
		s = string(rune('A'+n%26)) + s
		n /= 26
	}
	return s
}

// sheetRowRange builds an A1 range for a single row starting at column A,
// e.g. sheetRowRange("ForecastSetup", 5, 124) -> "'ForecastSetup'!A5:DT5".
func sheetRowRange(sheetName string, rowNum, numCols int) string {
	return fmt.Sprintf("'%s'!A%d:%s%d", sheetName, rowNum, colLetter(numCols), rowNum)
}

// ensureSheetWithHeader creates the sheet tab and writes its bold header row
// the first time it's used, mirroring the `if (!sheet) { ... }` blocks in
// Code.gs's saveRevenueForecast/saveSetup.
func ensureSheetWithHeader(ctx context.Context, svc *sheets.Service, sheetID, sheetName string, header []interface{}) error {
	meta, err := svc.Spreadsheets.Get(sheetID).Fields("sheets.properties").Context(ctx).Do()
	if err != nil {
		return err
	}
	for _, sh := range meta.Sheets {
		if sh.Properties.Title == sheetName {
			return nil
		}
	}

	_, err = svc.Spreadsheets.BatchUpdate(sheetID, &sheets.BatchUpdateSpreadsheetRequest{
		Requests: []*sheets.Request{{
			AddSheet: &sheets.AddSheetRequest{
				Properties: &sheets.SheetProperties{Title: sheetName},
			},
		}},
	}).Context(ctx).Do()
	if err != nil {
		return err
	}

	rangeA1 := sheetRowRange(sheetName, 1, len(header))
	_, err = svc.Spreadsheets.Values.Update(sheetID, rangeA1, &sheets.ValueRange{Values: [][]interface{}{header}}).
		ValueInputOption("RAW").Context(ctx).Do()
	return err
}
