package platform

import "testing"

func TestCloneCheckResultCopiesDetails(t *testing.T) {
	original := CheckResult{Details: map[string]any{"owner": "cache"}}
	c := CloneCheckResult(original)
	c.Details["owner"] = "caller"
	if original.Details["owner"] != "cache" { t.Fatal("details were not copied") }
}

func TestCloneVerificationCopiesChecks(t *testing.T) {
	original := Verification{Checks: []CheckResult{{Name: "stable"}}}
	v := CloneVerification(original)
	v.Checks[0].Name = "caller"
	if original.Checks[0].Name != "stable" { t.Fatal("checks were not copied") }
}
