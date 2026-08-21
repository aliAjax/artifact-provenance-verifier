package application

import (
	"context"
	"github.com/example/artifact-provenance-verifier/internal/webhook/domain"
	"sync"
)

type Outbox struct {
	mu     sync.Mutex
	events []domain.Event
}

func NewOutbox() *Outbox { return &Outbox{events: []domain.Event{}} }
func (o *Outbox) Add(e domain.Event) {
	o.mu.Lock()
	defer o.mu.Unlock()
	o.events = append(o.events, e)
}
func (o *Outbox) Drain(_ context.Context) []domain.Event {
	o.mu.Lock()
	defer o.mu.Unlock()
	v := append([]domain.Event(nil), o.events...)
	o.events = o.events[:0]
	return v
}
