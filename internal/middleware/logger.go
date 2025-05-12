package middleware

import (
	"net/http"
	"time"

	"mini-social-network-api/pkg/logger"
)

func RequestLoggerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		next.ServeHTTP(w, r)

		duration := time.Since(start)

		logger.Log.WithFields(map[string]interface{}{
			"method":   r.Method,
			"path":     r.URL.Path,
			"duration": duration.Milliseconds(),
			"agent":    r.UserAgent(),
		}).Info("Handled request")
	})
}
