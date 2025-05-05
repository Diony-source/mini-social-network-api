package middleware

import (
	"context"
	"net/http"
	"strings"

	"mini-social-network-api/pkg/auth"
)

type contextKey string

const ContextUserIDKey = contextKey("user_id")

func JWTAuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		header := r.Header.Get("Authorization")
		if header == "" || !strings.HasPrefix(header, "Bearer ") {
			http.Error(w, "missing or invalid authorization header", http.StatusUnauthorized)
			return
		}

		tokenStr := strings.TrimPrefix(header, "Bearer ")
		userID, err := auth.ParseToken(tokenStr)
		if err != nil {
			http.Error(w, "invalid token", http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), ContextUserIDKey, userID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
