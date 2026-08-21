package adapter

import (
	"context"
	"errors"
	"testing"
)

func TestGetArtifactHonorsCanceledContext(t *testing.T) {
	ctx := context.Background()
	b := NewBlobStore(t.TempDir())
	if err := b.Put(ctx, "artifact.json", []byte("{}")); err != nil {
		t.Fatal(err)
	}
	canceled, cancel := context.WithCancel(ctx)
	cancel()
	if _, err := b.Get(canceled, "artifact.json"); !errors.Is(err, context.Canceled) {
		t.Fatalf("Get error=%v", err)
	}
}

func TestPutArtifactHonorsCanceledContext(t *testing.T) {
	b := NewBlobStore(t.TempDir())
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	if err := b.Put(ctx, "write.json", []byte("{}")); !errors.Is(err, context.Canceled) {
		t.Fatalf("Put error=%v", err)
	}
}

func TestDeleteArtifactHonorsCanceledContext(t *testing.T) {
	b := NewBlobStore(t.TempDir())
	if err := b.Put(context.Background(), "delete.json", []byte("{}")); err != nil {
		t.Fatal(err)
	}
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	if err := b.Delete(ctx, "delete.json"); !errors.Is(err, context.Canceled) {
		t.Fatalf("Delete error=%v", err)
	}
}
