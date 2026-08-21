package domain

import "github.com/example/artifact-provenance-verifier/internal/platform"

type Entity = platform.Artifact

func IsImmutable(a platform.Artifact) bool { return a.Digest != "" }
