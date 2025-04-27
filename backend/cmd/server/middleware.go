package main

import (
	"net/http"
	"time"
)

func (app *application) withLogging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		// Call the next handler
		next.ServeHTTP(w, r)

		duration := time.Since(start)

		app.logger.Infow("incoming request",
			"method", r.Method,
			"path", r.URL.Path,
			"duration", duration,
			"remote", r.RemoteAddr,
			"user_agent", r.UserAgent(),
		)
	})
}
