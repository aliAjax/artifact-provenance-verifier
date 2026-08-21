package domain

import (
	"fmt"
	"github.com/example/artifact-provenance-verifier/internal/platform"
)

type Entity = platform.Artifact

func IsImmutable(a platform.Artifact) bool { return a.Digest != "" }
func ApplyStatus(a *Entity, next platform.Status) error {
	if a == nil {
		return fmt.Errorf("artifact is nil")
	}
	a.Status = string(next)
	return nil
}
