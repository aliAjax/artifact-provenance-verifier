package infrastructure

import (
	"crypto/sha256"
	"encoding/hex"
	"io"
)

func DigestReader(r io.Reader) (string, error) {
	h := sha256.New()
	if _, e := io.Copy(h, r); e != nil {
		return "", e
	}
	return "sha256:" + hex.EncodeToString(h.Sum(nil)), nil
}
