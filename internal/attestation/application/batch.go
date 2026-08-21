package application

import (
	"context"
	"github.com/example/artifact-provenance-verifier/internal/platform"
)

func (s *Service) SubmitBatch(ctx context.Context, items []platform.Attestation) ([]platform.Attestation, error) {
	out := make([]platform.Attestation, 0, len(items))
	for _, x := range items {
		v, e := s.Submit(ctx, x)
		if e != nil {
			out = append(out, platform.Attestation{ID: x.ID, ArtifactID: x.ArtifactID, Status: "rejected"})
			continue
		}
		out = append(out, v)
	}
	return out, nil
}
