// Package handler implements POST /api/export-excel, the Go replacement for
// Code.gs's exportFlashPL(): build a formatted Flash P&L report into a new
// Google Sheet (owned by the service account) and share it with the signed-in
// user so the returned .xlsx export link is actually openable by them.
package handler

import (
	"encoding/json"
	"net/http"

	"flashpnl/internal/apiutil"
	"flashpnl/internal/auth"
	"flashpnl/internal/export"
	"flashpnl/internal/sheetsdata"
)

type exportExcelRequest struct {
	Month  int           `json:"month"`
	Year   int           `json:"year"`
	PLData export.PLData `json:"plData"`
}

type exportExcelResponse struct {
	URL string `json:"url"`
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

	var req exportExcelRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		apiutil.WriteError(w, http.StatusBadRequest, err)
		return
	}

	sheetsSvc, err := sheetsdata.NewClient(ctx)
	if err != nil {
		apiutil.WriteError(w, http.StatusInternalServerError, err)
		return
	}
	sheetID, err := sheetsdata.SheetID()
	if err != nil {
		apiutil.WriteError(w, http.StatusInternalServerError, err)
		return
	}
	driveSvc, err := export.NewDriveClient(ctx)
	if err != nil {
		apiutil.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	origName, err := sheetsdata.SpreadsheetName(ctx, sheetsSvc, sheetID)
	if err != nil {
		apiutil.WriteError(w, http.StatusInternalServerError, err)
		return
	}
	tz, err := sheetsdata.TimeZone(ctx, sheetsSvc, sheetID)
	if err != nil {
		apiutil.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	url, err := export.Export(ctx, sheetsSvc, driveSvc, tz, origName, req.Month, req.Year, req.PLData, email)
	if err != nil {
		apiutil.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	apiutil.WriteJSON(w, http.StatusOK, exportExcelResponse{URL: url})
}
