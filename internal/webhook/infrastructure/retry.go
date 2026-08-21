package infrastructure

import (
	"github.com/example/artifact-provenance-verifier/internal/webhook/domain"
	"time"
)

func Backoff(attempt int) time.Duration {
	if attempt < 1 {
		attempt = 1
	}
	if attempt > 6 {
		attempt = 6
	}
	return time.Duration(1<<attempt) * 100 * time.Millisecond
}

func CopyEvents(in []domain.Event) []domain.Event { return append([]domain.Event(nil), in...) }
