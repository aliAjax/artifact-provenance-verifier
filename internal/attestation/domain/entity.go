package domain

import "github.com/example/artifact-provenance-verifier/internal/platform"

type Entity = platform.Attestation

func Accepted(a platform.Attestation) bool { return a.Status == "accepted" }
