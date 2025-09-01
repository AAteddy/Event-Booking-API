package middleware

import (
	"net/http"

	"golang.org/x/time/rate"
)

func RateLimitMiddleware(limit float64, burst int) func(http.Handler) http.Handler {
	limiter := rate.NewLimiter(rate.Limit(limit), burst)
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()
			if err := limiter.Wait(ctx); err != nil {
				http.Error(w, "Too Many Requests", http.StatusTooManyRequests)
				return
			}
			next.ServeHTTP(w, r)
		})
	}
}
