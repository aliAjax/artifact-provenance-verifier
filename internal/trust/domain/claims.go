package domain

import (
	"github.com/example/artifact-provenance-verifier/internal/platform"
	"strings"
)

func SubjectMatches(a platform.Attestation, digest string) bool {
	return strings.EqualFold(a.Subject, digest)
}
func TrustedType(a platform.Attestation, types []string) bool {
	for _, t := range types {
		if a.PredicateType == t {
			return true
		}
	}
	return false
}
