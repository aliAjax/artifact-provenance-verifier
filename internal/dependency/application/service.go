package application

import (
	"context"
	"fmt"
	"github.com/example/artifact-provenance-verifier/internal/platform"
)

type Service struct{}

func New() *Service { return &Service{} }
func (*Service) Validate(s platform.SBOM) error {
	for _, e := range s.Dependencies {
		if e.From == "" || e.To == "" {
			return fmt.Errorf("dependency endpoints required")
		}
	}
	return nil
}
func (*Service) Find(_ context.Context, s platform.SBOM, purl string) []platform.Component {
	out := []platform.Component{}
	for _, c := range s.Components {
		if c.PURL == purl {
			out = append(out, c)
		}
	}
	return out
}
