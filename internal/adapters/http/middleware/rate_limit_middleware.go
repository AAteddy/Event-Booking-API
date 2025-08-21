package middleware

import "net/http"

func RateLimitMiddleware(next http.Handler) http.Handler {
	return nil
}
