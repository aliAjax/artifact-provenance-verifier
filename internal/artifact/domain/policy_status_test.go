package domain

import (
	"github.com/example/artifact-provenance-verifier/internal/platform"
	"testing"
)

func TestCanRestoreRequiresWithdrawn(t *testing.T) {
	if CanRestore(platform.StatusActive) || CanRestore(platform.StatusIsolated) || CanRestore(platform.StatusRestored) {
		t.Fatal("non-withdrawn artifact was considered restorable")
	}
	if !CanRestore(platform.StatusWithdrawn) {
		t.Fatal("withdrawn artifact was not considered restorable")
	}
}
