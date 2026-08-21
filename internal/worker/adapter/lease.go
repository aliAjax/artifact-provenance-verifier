package adapter

import (
	"context"
	"sync"
)

type LeaseStore struct {
	mu     sync.Mutex
	owners map[string]string
}

func New() *LeaseStore { return &LeaseStore{owners: map[string]string{}} }
func (l *LeaseStore) Acquire(_ context.Context, id, owner string) bool {
	l.mu.Lock()
	defer l.mu.Unlock()
	if _, ok := l.owners[id]; ok {
		return false
	}
	l.owners[id] = owner
	return true
}
func (l *LeaseStore) Release(_ context.Context, id string) {
	l.mu.Lock()
	defer l.mu.Unlock()
	delete(l.owners, id)
}
