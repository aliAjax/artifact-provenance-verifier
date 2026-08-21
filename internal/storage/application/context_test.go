package application

import (
	"context"
	"errors"
	"testing"
)

func TestArtifactRejectsCanceledContextBeforeRepository(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	var s *Service
	_, err := s.Artifact(ctx, "artifact")
	if !errors.Is(err, context.Canceled) {
		t.Fatalf("Artifact error=%v", err)
	}
}

func TestArtifactsRejectsCanceledContextBeforeRepository(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	var s *Service
	_, err := s.Artifacts(ctx, "artifact")
	if !errors.Is(err, context.Canceled) {
		t.Fatalf("Artifacts error=%v", err)
	}
}
