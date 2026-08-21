package infrastructure

import (
	"github.com/example/artifact-provenance-verifier/internal/platform"
	"sync"
)

type Cache struct {
	mu    sync.RWMutex
	items map[string]platform.Verification
}

func NewCache() *Cache { return &Cache{items: map[string]platform.Verification{}} }
func (c *Cache) Get(k string) (platform.Verification, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	v, ok := c.items[k]
	return v, ok
}
func (c *Cache) Put(k string, v platform.Verification) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.items[k] = v
}
