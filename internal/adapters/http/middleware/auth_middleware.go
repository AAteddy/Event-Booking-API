package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/AAteddy/event-booking-api/internal/infrastructure/cache"
	"github.com/golang-jwt/jwt/v5"
)

type contextKey string

const UserIDKey contextKey = "user_id"

// AuthMiddleware is a middleware that checks for a valid JWT token in the Authorization header
// and sets the user ID in the request context. It uses the AuthRepository to validate the token.
// If the token is valid, the next handler is called. If the token is invalid, an error is returned.
func AuthMiddleware(authRepo *cache.AuthRepositoryImpl, jwtSecret string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			authHeader := r.Header.Get("Authorization")
			if authHeader == "" {
				http.Error(w, "Authorization header is required", http.StatusUnauthorized)
				return
			}

			parts := strings.SplitN(authHeader, " ", 2)
			if len(parts) != 2 || parts[0] != "Bearer" {
				http.Error(w, "Authorization header format must be Bearer {token}", http.StatusUnauthorized)
				return
			}

			tokenStr := parts[1]
			isBlacklisted, err := authRepo.IsTokenBlacklisted(tokenStr)
			if err != nil {
				http.Error(w, "Error validating token", http.StatusInternalServerError)
				return
			}
			if isBlacklisted {
				http.Error(w, "Token is Blacklisted", http.StatusUnauthorized)
				return
			}

			token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, http.ErrAbortHandler
				}
				return []byte(jwtSecret), nil
			})
			if err != nil || !token.Valid {
				http.Error(w, "Invalid token", http.StatusUnauthorized)
				return
			}

			claims, ok := token.Claims.(jwt.MapClaims)
			if !ok {
				http.Error(w, "Invalid token claims", http.StatusUnauthorized)
				return
			}

			userID, ok := claims["sub"].(string)
			if !ok {
				http.Error(w, "Invalid user ID in token", http.StatusUnauthorized)
				return
			}

			ctx := context.WithValue(r.Context(), UserIDKey, userID)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
