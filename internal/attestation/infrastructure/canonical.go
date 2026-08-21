package infrastructure

import (
	"crypto/sha256"
	"encoding/hex"
	"github.com/example/artifact-provenance-verifier/internal/platform"
)

func CanonicalDigest(a platform.Attestation) string {
	b, _ := platform.MarshalCanonical(a.Predicate)
	h := sha256.Sum256(b)
	return "sha256:" + hex.EncodeToString(h[:])
}
