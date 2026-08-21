package application

import (
	"context"
	"github.com/example/artifact-provenance-verifier/internal/platform"
)

type Difference struct{ Added, Removed []platform.Component }

func (s *Service) Diff(_ context.Context, a, b platform.SBOM) Difference {
	am := map[string]platform.Component{}
	bm := map[string]platform.Component{}
	for _, c := range a.Components {
		am[c.Name+"@"+c.Version] = c
	}
	for _, c := range b.Components {
		bm[c.Name+"@"+c.Version] = c
	}
	d := Difference{}
	for k, c := range bm {
		if _, ok := am[k]; !ok {
			d.Added = append(d.Added, c)
		}
	}
	for k, c := range am {
		if _, ok := bm[k]; !ok {
			d.Removed = append(d.Removed, c)
		}
	}
	return d
}
