// Package handler implements GET /api/config: exposes the (non-secret)
// Google OAuth client ID to the frontend so it isn't hardcoded into
// index.html and can be rotated without a frontend redeploy.
package handler

import (
	"net/http"
	"os"

	"flashpnl/internal/apiutil"
)

type configResponse struct {
	GoogleClientID string `json:"googleClientId"`
}

func Handler(w http.ResponseWriter, r *http.Request) {
	apiutil.WriteJSON(w, http.StatusOK, configResponse{
		GoogleClientID: os.Getenv("GOOGLE_OAUTH_CLIENT_ID"),
	})
}
