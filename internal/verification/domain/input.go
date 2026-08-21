package domain

import (
	"fmt"
	"github.com/example/artifact-provenance-verifier/internal/platform"
)

type Input struct {
	ArtifactID      string
	Attestations    []platform.Attestation
	SBOM            *platform.SBOM
	Vulnerabilities []platform.Vulnerability
}

func (i Input) Validate() error {
	if i.ArtifactID == "" {
		return fmt.Errorf("artifact id required")
	}
	if i.SBOM == nil && len(i.Vulnerabilities) > 0 {
		return fmt.Errorf("vulnerabilities require sbom")
	}
	return nil
}
func (i Input) Digest() string { d, _ := platform.HashJSON(i); return d }
