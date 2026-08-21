package infrastructure

import "github.com/example/artifact-provenance-verifier/internal/platform"

func HasCycle(s platform.SBOM) bool {
	adj := map[string][]string{}
	for _, e := range s.Dependencies {
		adj[e.From] = append(adj[e.From], e.To)
	}
	vis := map[string]bool{}
	stack := map[string]bool{}
	var dfs func(string) bool
	dfs = func(n string) bool {
		if stack[n] {
			return true
		}
		if vis[n] {
			return false
		}
		vis[n] = true
		stack[n] = true
		for _, x := range adj[n] {
			if dfs(x) {
				return true
			}
		}
		stack[n] = false
		return false
	}
	for n := range adj {
		if dfs(n) {
			return true
		}
	}
	return false
}
