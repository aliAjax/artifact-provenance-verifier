package adapter

import (
	"context"
	"errors"
	"testing"
)

func TestSQLStoreWithinPreservesContext(t *testing.T) {
	store := NewSQLStore("memory://")
	want := errors.New("operation failed")
	ctx := context.WithValue(context.Background(), struct{}{}, "request")
	if err := store.Within(ctx, func(got context.Context) error {
		if got.Value(struct{}{}) != "request" {
			t.Fatal("transaction context was replaced")
		}
		return want
	}); !errors.Is(err, want) {
		t.Fatalf("Within error=%v", err)
	}
}

func TestSQLStoreWithinRejectsNilReceiver(t *testing.T) {
	var store *SQLStore
	if err := store.Within(context.Background(), func(context.Context) error { return nil }); err == nil {
		t.Fatal("nil SQL store was accepted")
	}
}
