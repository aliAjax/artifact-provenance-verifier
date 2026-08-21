package api

import (
	"context"
	"net/http"
	"time"
)

type Dependency interface{ Ping(context.Context) error }

func DependencyHealth(dep []Dependency) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx, cancel := context.WithTimeout(r.Context(), 2*time.Second)
		defer cancel()
		for _, d := range dep {
			if e := d.Ping(ctx); e != nil {
				encode(w, 503, map[string]any{"status": "degraded", "error": e.Error()})
				return
			}
		}
		encode(w, 200, map[string]any{"status": "ok", "dependencies": len(dep)})
	}
}
