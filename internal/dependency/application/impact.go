package application

import (
	"context"
	"github.com/example/artifact-provenance-verifier/internal/platform"
)

type Impact struct {
	Component  platform.Component
	Dependents []string
}

func (s *Service) ImpactAnalysis(_ context.Context, sbom platform.SBOM, name string) Impact {
	var c platform.Component
	for _, x := range sbom.Components {
		if x.Name == name {
			c = x
		}
	}
	out := []string{}
	for _, e := range sbom.Dependencies {
		if e.To == name || e.To == c.Name+"@"+c.Version {
			out = append(out, e.From)
		}
	}
	return Impact{Component: c, Dependents: out}
}
