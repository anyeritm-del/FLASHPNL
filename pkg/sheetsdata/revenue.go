package sheetsdata

import (
	"context"
	"time"

	"google.golang.org/api/sheets/v4"
)

const revenueSheetName = "ForecastRevenue"

var revenueHeader = []interface{}{
	"Year", "Month", "SavedBy", "SavedAt",
	"Room", "Breakfast", "RestoFood", "RestoBev",
	"TapasFood", "TapasBev", "RSFood", "RSBev",
	"BQT", "Spa", "Laundry", "OtherIncome",
}

func rowToSavedRevenue(row []interface{}) *SavedRevenue {
	return &SavedRevenue{
		SavedBy: toStr(at(row, 2)),
		SavedAt: toStr(at(row, 3)),
		Rev: Revenue{
			Room: toFloat(at(row, 4)), Breakfast: toFloat(at(row, 5)),
			RestoF: toFloat(at(row, 6)), RestoB: toFloat(at(row, 7)),
			TapasF: toFloat(at(row, 8)), TapasB: toFloat(at(row, 9)),
			RsF: toFloat(at(row, 10)), RsB: toFloat(at(row, 11)),
			Bqt: toFloat(at(row, 12)), Spa: toFloat(at(row, 13)),
			Laundry: toFloat(at(row, 14)), OtherIncome: toFloat(at(row, 15)),
		},
	}
}

// GetRevenue mirrors getRevenue_(): returns nil (no error) if the sheet
// doesn't exist yet or no row matches the given period.
func GetRevenue(ctx context.Context, svc *sheets.Service, sheetID string, month, year int) (*SavedRevenue, error) {
	resp, err := svc.Spreadsheets.Values.Get(sheetID, revenueSheetName).
		ValueRenderOption("UNFORMATTED_VALUE").Context(ctx).Do()
	if err != nil {
		if isMissingSheetErr(err) {
			return nil, nil
		}
		return nil, err
	}
	for i, row := range resp.Values {
		if i == 0 {
			continue
		}
		if toInt(at(row, 0)) == year && toInt(at(row, 1)) == month {
			return rowToSavedRevenue(row), nil
		}
	}
	return nil, nil
}

// SaveRevenueForecast mirrors saveRevenueForecast(): upserts the row for
// (year, month) and returns the formatted save timestamp.
func SaveRevenueForecast(ctx context.Context, svc *sheets.Service, sheetID string, month, year int, rev Revenue, savedByEmail string) (string, error) {
	if err := ensureSheetWithHeader(ctx, svc, sheetID, revenueSheetName, revenueHeader); err != nil {
		return "", err
	}

	tz, err := TimeZone(ctx, svc, sheetID)
	if err != nil {
		return "", err
	}
	loc, err := time.LoadLocation(tz)
	if err != nil {
		loc = time.UTC
	}
	timestamp := time.Now().In(loc).Format("02 Jan 2006 15:04")

	newRow := []interface{}{
		year, month, savedByEmail, timestamp,
		rev.Room, rev.Breakfast, rev.RestoF, rev.RestoB,
		rev.TapasF, rev.TapasB, rev.RsF, rev.RsB,
		rev.Bqt, rev.Spa, rev.Laundry, rev.OtherIncome,
	}

	resp, err := svc.Spreadsheets.Values.Get(sheetID, revenueSheetName).
		ValueRenderOption("UNFORMATTED_VALUE").Context(ctx).Do()
	if err != nil {
		return "", err
	}

	for i, row := range resp.Values {
		if i == 0 {
			continue
		}
		if toInt(at(row, 0)) == year && toInt(at(row, 1)) == month {
			rangeA1 := sheetRowRange(revenueSheetName, i+1, len(newRow))
			_, err := svc.Spreadsheets.Values.Update(sheetID, rangeA1, &sheets.ValueRange{Values: [][]interface{}{newRow}}).
				ValueInputOption("RAW").Context(ctx).Do()
			return timestamp, err
		}
	}

	_, err = svc.Spreadsheets.Values.Append(sheetID, revenueSheetName, &sheets.ValueRange{Values: [][]interface{}{newRow}}).
		ValueInputOption("RAW").Context(ctx).Do()
	return timestamp, err
}
