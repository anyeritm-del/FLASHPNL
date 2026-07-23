package sheetsdata

import (
	"context"
	"strings"

	"google.golang.org/api/sheets/v4"
)

// GetUserRole mirrors getUserRole_(): looks up email in the "Users" sheet
// (Email | Role | Name) and returns role "unknown" if not found.
func GetUserRole(ctx context.Context, svc *sheets.Service, sheetID, email string) (*UserInfo, error) {
	email = strings.ToLower(email)
	resp, err := svc.Spreadsheets.Values.Get(sheetID, "Users").
		ValueRenderOption("UNFORMATTED_VALUE").Context(ctx).Do()
	if err != nil {
		return nil, err
	}
	for i, row := range resp.Values {
		if i == 0 {
			continue
		}
		rowEmail := strings.ToLower(toStr(at(row, 0)))
		if rowEmail != "" && rowEmail == email {
			name := toStr(at(row, 2))
			if name == "" {
				name = email
			}
			return &UserInfo{
				Email: email,
				Role:  strings.ToLower(toStr(at(row, 1))),
				Name:  name,
			}, nil
		}
	}
	return &UserInfo{Email: email, Role: "unknown", Name: email}, nil
}
