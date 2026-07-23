package sheetsdata

import (
	"context"
	"errors"
	"strings"

	"google.golang.org/api/sheets/v4"
)

// ErrUnregistered is returned by RequireUser when the signed-in email has no
// row in the "Users" sheet — Google Sign-In only proves identity, not that
// this person should see the hotel's financial data.
var ErrUnregistered = errors.New("email tidak terdaftar, hubungi admin untuk mendapatkan akses")

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

// RequireUser looks up email like GetUserRole but rejects it with
// ErrUnregistered instead of returning a role of "unknown" — use this at
// every data-serving endpoint so Google Sign-In alone never grants access.
func RequireUser(ctx context.Context, svc *sheets.Service, sheetID, email string) (*UserInfo, error) {
	info, err := GetUserRole(ctx, svc, sheetID, email)
	if err != nil {
		return nil, err
	}
	if info.Role == "unknown" {
		return nil, ErrUnregistered
	}
	return info, nil
}
