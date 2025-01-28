package middleware

import (
	"log/slog"
	"net/http"
	"time"

	log_middleware "github.com/heisenberg8055/gotiny/internal/log"
)

func Middleware(next func(http.ResponseWriter, *http.Request), logger *slog.Logger) func(http.ResponseWriter, *http.Request) {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		timeStart := time.Now()
		next(w, r)
		timeElapsed := time.Since(timeStart)
		log_middleware.LogInfo(log_middleware.Response{Method: r.Method, Url: r.URL.Path, TimeTaken: timeElapsed.String()}, logger, "Request Received")
	})
}
