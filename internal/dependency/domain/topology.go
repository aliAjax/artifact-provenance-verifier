package domain

import "github.com/example/artifact-provenance-verifier/internal/platform"

func Roots(s platform.SBOM) []string {
	incoming := map[string]bool{}
	for _, e := range s.Dependencies {
		incoming[e.To] = true
	}
	out := []string{}
	for _, c := range s.Components {
		k := c.Name + "@" + c.Version
		if !incoming[k] {
			out = append(out, k)
		}
	}
	return out
}
func Outgoing(s platform.SBOM, node string) []string {
	out := []string{}
	for _, e := range s.Dependencies {
		if e.From == node {
			out = append(out, e.To)
		}
	}
	return out
}
