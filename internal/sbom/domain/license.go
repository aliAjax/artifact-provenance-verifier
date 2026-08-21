package domain

import (
	"github.com/example/artifact-provenance-verifier/internal/platform"
	"strings"
)

func LicenseSet(s platform.SBOM) map[string]bool {
	out := map[string]bool{}
	for _, c := range s.Components {
		if c.License != "" {
			out[strings.ToUpper(c.License)] = true
		}
	}
	return out
}
func DisallowedLicense(s platform.SBOM, blocked []string) string {
	set := LicenseSet(s)
	for _, v := range blocked {
		if set[strings.ToUpper(v)] {
			return v
		}
	}
	return ""
}
