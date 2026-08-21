package infrastructure

import (
	"encoding/json"
	"github.com/example/artifact-provenance-verifier/internal/platform"
	"sync"
)

type SnapshotStore struct {
	mu   sync.RWMutex
	data map[string][]byte
}

func NewSnapshotStore() *SnapshotStore { return &SnapshotStore{data: map[string][]byte{}} }
func (s *SnapshotStore) Save(k string, v any) error {
	b, e := json.Marshal(v)
	if e != nil {
		return e
	}
	s.mu.Lock()
	defer s.mu.Unlock()
	s.data[k] = append([]byte(nil), b...)
	return nil
}
func (s *SnapshotStore) Load(k string, v any) error {
	s.mu.RLock()
	b, ok := s.data[k]
	s.mu.RUnlock()
	if !ok {
		return platform.ErrNotFound
	}
	return json.Unmarshal(b, v)
}
