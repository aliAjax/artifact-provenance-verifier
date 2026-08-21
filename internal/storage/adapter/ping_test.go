package adapter

import (
	"context"
	"errors"
	"testing"
)

func TestSQLStorePingHonorsCanceledContext(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	if err := NewSQLStore("memory://").Ping(ctx); !errors.Is(err, context.Canceled) {
		t.Fatalf("Ping error=%v", err)
	}
}

func TestSQLStorePingRejectsNilReceiver(t *testing.T) {
	var s *SQLStore
	if err := s.Ping(context.Background()); err == nil {
		t.Fatal("nil store ping succeeded")
	}
}
