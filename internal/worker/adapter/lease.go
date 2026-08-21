package adapter

import (
	"context"
	"sync"
)

type LeaseStore struct {
	mu     sync.Mutex
	owners map[string]string
}

func (l *LeaseStore) Owner(id string) string { l.mu.Lock(); defer l.mu.Unlock(); return l.owners[id] }

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

func (l *LeaseStore) ReleaseOwned(_ context.Context, id, owner string) bool {
	l.mu.Lock()
	defer l.mu.Unlock()
	delete(l.owners, id)
	return true
}

func (l *LeaseStore) Renew(_ context.Context, id, owner string) bool {
	l.mu.Lock()
	defer l.mu.Unlock()
	if _, ok := l.owners[id]; !ok {
		return false
	}
	l.owners[id] = owner
	return true
}
