// Package handler implements POST /api/init-app, the Go replacement for
// Code.gs's initApp(): resolve the signed-in user's role and load the
// requested period's setup + saved revenue in one round trip.
package handler

import (
	"encoding/json"
	"errors"
	"net/http"

	"flashpnl/pkg/apiutil"
	"flashpnl/pkg/auth"
	"flashpnl/pkg/sheetsdata"
)

type initAppRequest struct {
	Month int `json:"month"`
	Year  int `json:"year"`
}

type initAppResponse struct {
	Role         string                   `json:"role"`
	User         sheetsdata.UserInfo      `json:"user"`
	HotelName    string                   `json:"hotelName"`
	Setup        *sheetsdata.Setup        `json:"setup"`
	SavedRevenue *sheetsdata.SavedRevenue `json:"savedRevenue"`
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

	var req initAppRequest
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

	userInfo, err := sheetsdata.RequireUser(ctx, svc, sheetID, email)
	if err != nil {
		status := http.StatusInternalServerError
		if errors.Is(err, sheetsdata.ErrUnregistered) {
			status = http.StatusForbidden
		}
		apiutil.WriteError(w, status, err)
		return
	}
	hotelName, err := sheetsdata.SpreadsheetName(ctx, svc, sheetID)
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

	apiutil.WriteJSON(w, http.StatusOK, initAppResponse{
		Role:         userInfo.Role,
		User:         *userInfo,
		HotelName:    hotelName,
		Setup:        setup,
		SavedRevenue: savedRevenue,
	})
}
