package application

import (
	"github.com/example/artifact-provenance-verifier/internal/platform"
	"testing"
)

func TestCloneVerificationForSummaryIsIndependent(t *testing.T) {
	original := platform.Verification{Checks: []platform.CheckResult{{Name: "stable", Details: map[string]any{"owner": "cache"}}}}
	copy := CloneVerificationForSummary(original)
	copy.Checks[0].Name = "caller"
	copy.Checks[0].Details["owner"] = "caller"
	if original.Checks[0].Name != "stable" || original.Checks[0].Details["owner"] != "cache" {
		t.Fatalf("summary clone aliases verification: %#v", original)
	}
}
