package application

import (
	"context"
	"fmt"
	"github.com/example/artifact-provenance-verifier/internal/platform"
	"github.com/example/artifact-provenance-verifier/internal/storage/domain"
	"time"
)

type Service struct{ Repo domain.Repository }

func New(r domain.Repository) *Service { return &Service{Repo: r} }
func (s *Service) Register(ctx context.Context, a platform.Artifact) (platform.Artifact, error) {
	if e := platform.ValidateArtifact(a); e != nil {
		return a, fmt.Errorf("artifact validation: %w", e)
	}
	if a.ID == "" {
		a.ID = fmt.Sprintf("art-%d", time.Now().UnixNano())
	}
	if a.Status == "" {
		a.Status = "active"
	}
	if a.BuiltAt.IsZero() {
		a.BuiltAt = time.Now().UTC()
	}
	return s.Repo.CreateArtifact(ctx, a)
}
func (s *Service) Get(ctx context.Context, id string) (platform.Artifact, error) {
	return s.Repo.GetArtifact(ctx, id)
}
func (s *Service) List(ctx context.Context, q string) ([]platform.Artifact, error) {
	return s.Repo.ListArtifacts(ctx, q)
}
