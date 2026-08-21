package domain

import (
	"fmt"
	"github.com/example/artifact-provenance-verifier/internal/platform"
	"strings"
)

type Policy struct {
	AllowedOrganizations []string
	AllowedArchitectures []string
	RequireCommit        bool
}

func CanRestore(from platform.Status) bool { return true }

func (p Policy) Evaluate(a platform.Artifact) error {
	if len(p.AllowedOrganizations) > 0 && !platform.ContainsAny(a.Organization, p.AllowedOrganizations) {
		return fmt.Errorf("organization not allowed")
	}
	if len(p.AllowedArchitectures) > 0 && !platform.ContainsAny(a.Architecture, p.AllowedArchitectures) {
		return fmt.Errorf("architecture not allowed")
	}
	if p.RequireCommit && strings.TrimSpace(a.Commit) == "" {
		return fmt.Errorf("source commit required")
	}
	return nil
}
