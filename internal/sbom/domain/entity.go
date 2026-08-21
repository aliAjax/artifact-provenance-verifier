package domain

import "github.com/example/artifact-provenance-verifier/internal/platform"

type Entity = platform.SBOM

func ComponentCount(s platform.SBOM) int { return len(s.Components) }
