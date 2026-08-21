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
func (s *Service) Submit(ctx context.Context, a platform.Attestation) (platform.Attestation, error) {
	if a.ArtifactID == "" || a.BodyDigest == "" {
		return a, fmt.Errorf("artifact_id and body_digest required")
	}
	if a.ID == "" {
		a.ID = fmt.Sprintf("att-%d", time.Now().UnixNano())
	}
	if a.CreatedAt.IsZero() {
		a.CreatedAt = time.Now().UTC()
	}
	a.Status = "accepted"
	if a.Type == "" {
		a.Type = "in-toto"
	}
	if a.Predicate == nil {
		a.Predicate = map[string]any{}
	}
	if err := verifyClaim(a); err != nil {
		a.Status = "rejected"
		return a, err
	}
	return s.Repo.PutAttestation(ctx, a)
}
func verifyClaim(a platform.Attestation) error {
	if a.Subject == "" {
		return fmt.Errorf("subject is required")
	}
	if a.PredicateType == "" {
		return fmt.Errorf("predicate_type is required")
	}
	return nil
}
func (s *Service) List(ctx context.Context, id string) ([]platform.Attestation, error) {
	return s.Repo.ListAttestations(ctx, id)
}
