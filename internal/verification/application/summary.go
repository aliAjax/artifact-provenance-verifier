package application

import (
	"context"
	"github.com/example/artifact-provenance-verifier/internal/platform"
)

func (s *Service) Summary(ctx context.Context, id, digest string) (map[string]any, error) {
	v, e := s.Repo.FindVerification(ctx, id, digest)
	if e != nil {
		return nil, e
	}
	passed := 0
	failed := 0
	for _, c := range v.Checks {
		if c.Status == "pass" {
			passed++
		}
		if c.Status == "fail" {
			failed++
		}
	}
	return map[string]any{"id": v.ID, "conclusion": v.Conclusion, "passed": passed, "failed": failed, "reused": v.Reused}, nil
}

func CacheSummaryChecks(v platform.Verification) int { return len(v.Checks) }
