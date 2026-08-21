package domain

import (
	"github.com/example/artifact-provenance-verifier/internal/platform"
	"testing"
)

func TestApplyStatusRejectsInvalidTarget(t *testing.T) {
	a := &Entity{Status: string(platform.StatusActive)}
	if err := ApplyStatus(a, platform.Status("unknown")); err == nil {
		t.Fatal("domain accepted an invalid status target")
	}
}
