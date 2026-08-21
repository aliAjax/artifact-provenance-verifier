package domain

import (
	"fmt"
	"github.com/example/artifact-provenance-verifier/internal/platform"
	"strings"
)

func ValidatePredicate(a platform.Attestation) error {
	if !strings.Contains(a.PredicateType, "/") {
		return fmt.Errorf("predicate type must be URI")
	}
	if len(a.Predicate) > 128 {
		return fmt.Errorf("predicate has too many fields")
	}
	return nil
}
func IsSLSA(a platform.Attestation) bool {
	return strings.Contains(strings.ToLower(a.PredicateType), "slsa")
}
