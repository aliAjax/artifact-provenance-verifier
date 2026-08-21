package infrastructure

import (
	"github.com/example/artifact-provenance-verifier/internal/platform"
	"regexp"
)

var coordinate = regexp.MustCompile(`^[a-zA-Z0-9._/-]+$`)

func ValidCoordinate(a platform.Artifact) bool {
	return coordinate.MatchString(a.Name) && coordinate.MatchString(a.Repository)
}
