package application

import (
	"context"
	"github.com/example/artifact-provenance-verifier/internal/platform"
	"github.com/example/artifact-provenance-verifier/internal/storage/domain"
)

type Service struct{ Repo domain.Repository }

func NewService(r domain.Repository) *Service { return &Service{Repo: r} }
func (s *Service) Artifact(ctx context.Context, id string) (platform.Artifact, error) {
	return s.Repo.GetArtifact(ctx, id)
}
func (s *Service) Artifacts(ctx context.Context, q string) ([]platform.Artifact, error) {
	return s.Repo.ListArtifacts(ctx, q)
}
