package infrastructure

import (
	"testing"
	"time"
)

func TestNilCircuitReceiverIsClosed(t *testing.T) {
	var c *Circuit
	if c.Allow(time.Now()) {
		t.Fatal("nil circuit was treated as open")
	}
	c.RecordFailure(time.Now())
	c.RecordSuccess()
}
