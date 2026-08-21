package domain

import "github.com/example/artifact-provenance-verifier/internal/platform"

func Decide(checks []platform.CheckResult) string {
	for _, c := range checks {
		if c.Status == "fail" && c.Severity != "low" {
			return "fail"
		}
	}
	return "pass"
}
func Failed(checks []platform.CheckResult) []platform.CheckResult {
	out := []platform.CheckResult{}
	for _, c := range checks {
		if c.Status == "fail" {
			out = append(out, c)
		}
	}
	return out
}
