package infrastructure

import (
	"crypto/ed25519"
	"encoding/base64"
)

func VerifyEd25519(pub ed25519.PublicKey, msg []byte, sig string) bool {
	b, e := base64.StdEncoding.DecodeString(sig)
	return e == nil && ed25519.Verify(pub, msg, b)
}
