package adapter

import (
	"encoding/json"
	"fmt"
	"github.com/example/artifact-provenance-verifier/internal/platform"
)

func Parse(raw []byte, format string) (platform.SBOM, error) {
	var x struct {
		Components   []platform.Component      `json:"components"`
		Dependencies []platform.DependencyEdge `json:"dependencies"`
		SPDXID       string                    `json:"spdxVersion"`
		Serial       string                    `json:"serialNumber"`
	}
	if err := json.Unmarshal(raw, &x); err != nil {
		return platform.SBOM{}, fmt.Errorf("parse sbom: %w", err)
	}
	return platform.SBOM{Format: format, Components: x.Components, Dependencies: x.Dependencies, Serial: x.Serial}, nil
}
