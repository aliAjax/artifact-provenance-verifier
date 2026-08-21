package application

import (
	"context"
	"github.com/example/artifact-provenance-verifier/internal/platform"
)

func (s *Service) VerifySubject(_ context.Context, a platform.Attestation, digest string) bool {
	return a.Subject == digest && a.Status == "accepted"
}
