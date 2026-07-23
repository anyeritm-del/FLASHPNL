// Package handler implements POST /api/save-revenue, the Go replacement for
// Code.gs's saveRevenueForecast(): upsert the revenue row for a period. Any
// registered user may save revenue, matching the original's lack of a role
// check here (only saveSetup was accounting-only) — but the email must still
// exist in the Users sheet.
package handler

import (
	"encoding/json"
	"errors"
	"net/http"

	"flashpnl/pkg/apiutil"
	"flashpnl/pkg/auth"
	"flashpnl/pkg/sheetsdata"
)

type saveRevenueRequest struct {
	Month   int                `json:"month"`
	Year    int                `json:"year"`
	RevData sheetsdata.Revenue `json:"revData"`
}

type saveRevenueResponse struct {
	SavedAt string `json:"savedAt"`
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

	var req saveRevenueRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		apiutil.WriteError(w, http.StatusBadRequest, err)
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

	if _, err := sheetsdata.RequireUser(ctx, svc, sheetID, email); err != nil {
		status := http.StatusInternalServerError
		if errors.Is(err, sheetsdata.ErrUnregistered) {
			status = http.StatusForbidden
		}
		apiutil.WriteError(w, status, err)
		return
	}

	savedAt, err := sheetsdata.SaveRevenueForecast(ctx, svc, sheetID, req.Month, req.Year, req.RevData, email)
	if err != nil {
		apiutil.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	apiutil.WriteJSON(w, http.StatusOK, saveRevenueResponse{SavedAt: savedAt})
}
