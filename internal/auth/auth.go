// Package auth verifies Google Sign-In ID tokens, replacing the identity
// Session.getActiveUser() used to provide for free inside Apps Script.
package auth

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"strings"

	"google.golang.org/api/idtoken"
)

var ErrNoToken = errors.New("missing Authorization header")

// VerifyIDToken extracts the "Authorization: Bearer <idToken>" header from r,
// verifies it against Google's public keys with audience GOOGLE_OAUTH_CLIENT_ID,
// and returns the signed-in user's email. No Workspace-domain restriction is
// applied — any verified Google account is accepted.
func VerifyIDToken(ctx context.Context, r *http.Request) (string, error) {
	authHeader := r.Header.Get("Authorization")
	token, ok := strings.CutPrefix(authHeader, "Bearer ")
	if !ok || token == "" {
		return "", ErrNoToken
	}

	clientID := os.Getenv("GOOGLE_OAUTH_CLIENT_ID")
	if clientID == "" {
		return "", fmt.Errorf("GOOGLE_OAUTH_CLIENT_ID is not set")
	}

	payload, err := idtoken.Validate(ctx, token, clientID)
	if err != nil {
		return "", fmt.Errorf("invalid ID token: %w", err)
	}

	email, _ := payload.Claims["email"].(string)
	if email == "" {
		return "", errors.New("ID token has no email claim")
	}
	return email, nil
}
