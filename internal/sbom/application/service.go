package application

import (
	"context"
	"fmt"
	"github.com/example/artifact-provenance-verifier/internal/platform"
	"github.com/example/artifact-provenance-verifier/internal/storage/domain"
	"strings"
	"time"
)

type Service struct{ Repo domain.Repository }

func New(r domain.Repository) *Service { return &Service{Repo: r} }
func (s *Service) Submit(ctx context.Context, b platform.SBOM) (platform.SBOM, error) {
	if b.ArtifactID == "" {
		return b, fmt.Errorf("artifact_id required")
	}
	if b.Format != "SPDX" && b.Format != "CycloneDX" {
		return b, fmt.Errorf("format must be SPDX or CycloneDX")
	}
	if len(b.Components) == 0 {
		return b, fmt.Errorf("sbom needs components")
	}
	seen := map[string]bool{}
	for _, c := range b.Components {
		if e := platform.ValidateComponent(c); e != nil {
			return b, e
		}
		k := strings.ToLower(c.Name + "@" + c.Version)
		if seen[k] {
			return b, fmt.Errorf("duplicate component %s", k)
		}
		seen[k] = true
	}
	if b.ID == "" {
		b.ID = fmt.Sprintf("sbom-%d", time.Now().UnixNano())
	}
	if b.CreatedAt.IsZero() {
		b.CreatedAt = time.Now().UTC()
	}
	if b.Digest == "" {
		d, _ := platform.HashJSON(b)
		b.Digest = d
	}
	return s.Repo.PutSBOM(ctx, b)
}
func (s *Service) Get(ctx context.Context, id string) (platform.SBOM, error) {
	return s.Repo.GetSBOM(ctx, id)
}
