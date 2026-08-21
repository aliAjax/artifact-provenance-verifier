package application

import (
	"context"
	"github.com/example/artifact-provenance-verifier/internal/platform"
	"github.com/example/artifact-provenance-verifier/internal/storage/infrastructure"
	"testing"
)

func TestLifecycleRejectsRestoredWithdrawn(t *testing.T) {
	ctx := context.Background()
	repo := infrastructure.NewMemory()
	a := platform.Artifact{ID: "art-status", Organization: "acme", Repository: "repo", Name: "artifact", Digest: "sha256:abcdef0123456789", Algorithm: "sha256", Status: string(platform.StatusRestored)}
	if _, err := repo.CreateArtifact(ctx, a); err != nil {
		t.Fatal(err)
	}
	if _, err := NewLifecycle(repo).SetStatus(ctx, a.ID, string(platform.StatusWithdrawn)); err == nil {
		t.Fatal("lifecycle accepted restored to withdrawn")
	}
}
