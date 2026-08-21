package application

import (
	"context"
	"testing"

	"github.com/example/artifact-provenance-verifier/internal/platform"
	"github.com/example/artifact-provenance-verifier/internal/storage/infrastructure"
)

func TestVerifyBlocksKnownHighVulnerability(t *testing.T) {
	ctx := context.Background()
	repo := infrastructure.NewMemory()
	a := platform.Artifact{ID: "art-1", Organization: "acme", Repository: "repo", Name: "api", Digest: "sha256:0123456789abcdef", Algorithm: "sha256"}
	if _, err := repo.CreateArtifact(ctx, a); err != nil {
		t.Fatal(err)
	}
	if _, err := repo.PutAttestation(ctx, platform.Attestation{ID: "att-1", ArtifactID: a.ID, BodyDigest: "sha256:abcdef0123456789", Status: "accepted"}); err != nil {
		t.Fatal(err)
	}
	if _, err := repo.PutSBOM(ctx, platform.SBOM{ID: "sbom-1", ArtifactID: a.ID, Digest: "sha256:1111111111111111", Components: []platform.Component{{Name: "openssl", Version: "3.0"}}}); err != nil {
		t.Fatal(err)
	}
	if _, err := repo.PutVulnerability(ctx, platform.Vulnerability{ID: "v-1", Identifier: "CVE-1", PackagePattern: "openssl", Severity: "high"}); err != nil {
		t.Fatal(err)
	}
	v, err := New(repo).Verify(ctx, Inputs{ArtifactID: a.ID, RequireAttestation: true, RequireSBOM: true, MaxSeverity: "high"})
	if err != nil {
		t.Fatal(err)
	}
	if v.Conclusion != "fail" {
		t.Fatalf("want fail, got %s", v.Conclusion)
	}
}

func TestVerificationSnapshotIsReused(t *testing.T) {
	ctx := context.Background()
	repo := infrastructure.NewMemory()
	a := platform.Artifact{ID: "art-2", Organization: "acme", Repository: "repo", Name: "cli", Digest: "sha256:2222222222222222", Algorithm: "sha256"}
	_, _ = repo.CreateArtifact(ctx, a)
	service := New(repo)
	in := Inputs{ArtifactID: a.ID}
	first, err := service.Verify(ctx, in)
	if err != nil {
		t.Fatal(err)
	}
	second, err := service.Verify(ctx, in)
	if err != nil {
		t.Fatal(err)
	}
	if first.ID != second.ID || !second.Reused {
		t.Fatalf("snapshot was not reused")
	}
}
