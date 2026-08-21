package domain

import "github.com/example/artifact-provenance-verifier/internal/platform"

type Rule interface {
	Name() string
	Evaluate(platform.Artifact, platform.SBOM, []platform.Attestation) platform.CheckResult
}
