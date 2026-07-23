// Package handler implements POST /api/save-setup, the Go replacement for
// Code.gs's saveSetup(): upsert the 124-column setup row for a period,
// restricted to the "accounting" role exactly like the original.
package handler

import (
	"encoding/json"
	"errors"
	"net/http"

	"flashpnl/internal/apiutil"
	"flashpnl/internal/auth"
	"flashpnl/internal/sheetsdata"
)

type saveSetupRequest struct {
	Month     int              `json:"month"`
	Year      int              `json:"year"`
	SetupData sheetsdata.Setup `json:"setupData"`
}

func Handler(w http.ResponseWriter, r *http.Request) {
	if !apiutil.RequirePost(w, r) {
		return
	}
	ctx := r.Context()

	email, err := auth.VerifyIDToken(ctx, r)
	if err != nil {
		apiutil.WriteError(w, http.StatusUnauthorized, err)
		return
	}

	svc, err := sheetsdata.NewClient(ctx)
	if err != nil {
		apiutil.WriteError(w, http.StatusInternalServerError, err)
		return
	}
	sheetID, err := sheetsdata.SheetID()
	if err != nil {
		apiutil.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	userInfo, err := sheetsdata.GetUserRole(ctx, svc, sheetID, email)
	if err != nil {
		apiutil.WriteError(w, http.StatusInternalServerError, err)
		return
	}
	if userInfo.Role != "accounting" {
		apiutil.WriteError(w, http.StatusForbidden, errors.New("Akses ditolak: hanya role Accounting yang dapat mengubah setup."))
		return
	}

	var req saveSetupRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		apiutil.WriteError(w, http.StatusBadRequest, err)
		return
	}

	if err := sheetsdata.SaveSetup(ctx, svc, sheetID, req.Month, req.Year, req.SetupData); err != nil {
		apiutil.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	apiutil.WriteJSON(w, http.StatusOK, map[string]bool{"ok": true})
}
