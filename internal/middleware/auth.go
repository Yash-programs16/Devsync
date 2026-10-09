package middleware

import (
	"context"
	"encoding/json"
	"net/http"
	"strings"

	"github.com/Yash-programs16/devsync/internal/token"
	"github.com/google/uuid"
)

type contextKey string

const userIDKey contextKey = "user_id"

// RequireAuth verifies the PASETO access token.
func RequireAuth(
	tokenMaker token.Maker,
) func(http.Handler) http.Handler {

	return func(next http.Handler) http.Handler {

		return http.HandlerFunc(func(
			w http.ResponseWriter,
			r *http.Request,
		) {

			authHeader := r.Header.Get("Authorization")

			if authHeader == "" {
				writeUnauthorized(w, "authorization header is required")
				return
			}

			parts := strings.Fields(authHeader)

			if len(parts) != 2 ||
				!strings.EqualFold(parts[0], "Bearer") {

				writeUnauthorized(w, "invalid authorization header")
				return
			}

			accessToken := parts[1]

			payload, err := tokenMaker.VerifyToken(accessToken)
			if err != nil {
				writeUnauthorized(w, "invalid or expired token")
				return
			}

			if payload == nil || payload.UserID == uuid.Nil {
				writeUnauthorized(w, "invalid token payload")
				return
			}

			ctx := context.WithValue(
				r.Context(),
				userIDKey,
				payload.UserID,
			)

			next.ServeHTTP(
				w,
				r.WithContext(ctx),
			)
		})
	}
}

// UserIDFromContext returns the authenticated user's ID.
func UserIDFromContext(
	ctx context.Context,
) (uuid.UUID, bool) {

	userID, ok := ctx.Value(userIDKey).(uuid.UUID)

	return userID, ok
}

func writeUnauthorized(
	w http.ResponseWriter,
	message string,
) {
	w.Header().Set(
		"Content-Type",
		"application/json",
	)

	w.WriteHeader(http.StatusUnauthorized)

	_ = json.NewEncoder(w).Encode(map[string]string{
		"error": message,
	})
}
