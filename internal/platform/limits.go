package platform

import (
	"fmt"
	"io"
	"time"
)

type Limits struct {
	MaxArtifactName     int
	MaxComponents       int
	MaxDependencyEdges  int
	MaxAttestationBytes int64
	MaxSBOMBytes        int64
	Timeout             time.Duration
}

func DefaultLimits() Limits {
	return Limits{MaxArtifactName: 256, MaxComponents: 100000, MaxDependencyEdges: 250000, MaxAttestationBytes: 2 << 20, MaxSBOMBytes: 10 << 20, Timeout: 10 * time.Second}
}
func CheckSize(n, max int64) error {
	if n < 0 {
		return fmt.Errorf("negative size")
	}
	if max > 0 && n > max {
		return fmt.Errorf("payload exceeds limit")
	}
	return nil
}
func ReadLimited(r io.Reader, max int64) ([]byte, error) {
	if e := CheckSize(max, max); e != nil {
		return nil, e
	}
	return io.ReadAll(io.LimitReader(r, max+1))
}
func ValidateLimits(l Limits) error {
	if l.MaxComponents <= 0 || l.MaxDependencyEdges <= 0 {
		return fmt.Errorf("limits must be positive")
	}
	if l.Timeout <= 0 {
		return fmt.Errorf("timeout must be positive")
	}
	return nil
}

func (l Limits) ArtifactOK(a Artifact) error {
	if len(a.Name) > l.MaxArtifactName {
		return fmt.Errorf("artifact name too long")
	}
	if e := ValidateTags(a.Tags); e != nil {
		return e
	}
	return nil
}

func (l Limits) SBOMOK(s SBOM) error {
	if len(s.Components) > l.MaxComponents {
		return fmt.Errorf("too many components")
	}
	if len(s.Dependencies) > l.MaxDependencyEdges {
		return fmt.Errorf("too many dependency edges")
	}
	return nil
}
