package domain

import "github.com/example/artifact-provenance-verifier/internal/platform"

type Graph struct {
	Nodes map[string]platform.Component
	Edges []platform.DependencyEdge
}

func Build(s platform.SBOM) Graph {
	g := Graph{Nodes: map[string]platform.Component{}, Edges: s.Dependencies}
	for _, c := range s.Components {
		g.Nodes[c.Name+"@"+c.Version] = c
	}
	return g
}
