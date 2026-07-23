// Package handler implements POST /api/load-period, the Go replacement for
// Code.gs's loadPeriodData(): fetch setup + saved revenue for a period the
// user switched to after initial load.
package handler

import (
	"encoding/json"
	"net/http"

	"flashpnl/pkg/apiutil"
	"flashpnl/pkg/auth"
	"flashpnl/pkg/sheetsdata"
)

type loadPeriodRequest struct {
	Month int `json:"month"`
	Year  int `json:"year"`
}

type loadPeriodResponse struct {
	Setup        *sheetsdata.Setup        `json:"setup"`
	SavedRevenue *sheetsdata.SavedRevenue `json:"savedRevenue"`
}

func Handler(w http.ResponseWriter, r *http.Request) {
	if !apiutil.RequirePost(w, r) {
		return
	}
	ctx := r.Context()

	if _, err := auth.VerifyIDToken(ctx, r); err != nil {
		apiutil.WriteError(w, http.StatusUnauthorized, err)
		return
	}

	var req loadPeriodRequest
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

	setup, err := sheetsdata.GetSetup(ctx, svc, sheetID, req.Month, req.Year)
	if err != nil {
		apiutil.WriteError(w, http.StatusInternalServerError, err)
		return
	}
	savedRevenue, err := sheetsdata.GetRevenue(ctx, svc, sheetID, req.Month, req.Year)
	if err != nil {
		apiutil.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	apiutil.WriteJSON(w, http.StatusOK, loadPeriodResponse{Setup: setup, SavedRevenue: savedRevenue})
}
