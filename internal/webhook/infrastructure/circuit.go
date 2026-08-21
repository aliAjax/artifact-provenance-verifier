package infrastructure

import (
	"sync"
	"time"
)

type Circuit struct {
	mu        sync.Mutex
	failures  int
	openUntil time.Time
}

func (c *Circuit) Allow(now time.Time) bool {
	c.mu.Lock()
	defer c.mu.Unlock()
	return now.After(c.openUntil)
}
func (c *Circuit) RecordFailure(now time.Time) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.failures++
	if c.failures >= 3 {
		c.openUntil = now.Add(time.Minute)
	}
}
func (c *Circuit) RecordSuccess() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.failures = 0
	c.openUntil = time.Time{}
}
