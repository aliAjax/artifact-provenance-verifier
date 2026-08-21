package domain

import (
	"crypto/sha256"
	"encoding/hex"
)

func EventID(typ, id string) string {
	h := sha256.Sum256([]byte(typ + ":" + id))
	return hex.EncodeToString(h[:])
}
