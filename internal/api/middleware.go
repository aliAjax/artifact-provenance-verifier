package api

import (
	"log/slog"
	"net/http"
	"time"
)

func RequestLogger(log *slog.Logger, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		rw := &capture{ResponseWriter: w, status: 200}
		next.ServeHTTP(rw, r)
		log.Info("http request", "method", r.Method, "path", r.URL.Path, "status", rw.status, "duration_ms", time.Since(start).Milliseconds())
	})
}

type capture struct {
	http.ResponseWriter
	status int
}

func (c *capture) WriteHeader(n int) { c.status = n; c.ResponseWriter.WriteHeader(n) }
func Recovery(log *slog.Logger, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if x := recover(); x != nil {
				log.Error("panic recovered", "error", x)
				writeErr(w, 500, http.ErrAbortHandler)
			}
		}()
		next.ServeHTTP(w, r)
	})
}
func ContentType(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			w.Header().Set("Content-Type", "application/json")
		}
		next.ServeHTTP(w, r)
	})
}
func CORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		if r.Method == http.MethodOptions {
			w.WriteHeader(204)
			return
		}
		next.ServeHTTP(w, r)
	})
}
