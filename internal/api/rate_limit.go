package api

import (
	"net/http"
	"sync"
	"time"
)

type Limiter struct {
	mu     sync.Mutex
	window time.Time
	count  int
	Max    int
}

func NewLimiter(max int) *Limiter { return &Limiter{Max: max, window: time.Now()} }
func (l *Limiter) Allow() bool {
	l.mu.Lock()
	defer l.mu.Unlock()
	now := time.Now()
	if now.Sub(l.window) >= time.Minute {
		l.window = now
		l.count = 0
	}
	if l.count >= l.Max {
		return false
	}
	l.count++
	return true
}
func RateLimit(l *Limiter, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !l.Allow() {
			w.Header().Set("Retry-After", "60")
			encode(w, 429, ErrorBody{Code: "rate_limited", Message: "too many requests"})
			return
		}
		next.ServeHTTP(w, r)
	})
}
