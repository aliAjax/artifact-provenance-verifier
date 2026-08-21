package infrastructure

import (
	"encoding/json"
	"fmt"
	"github.com/example/artifact-provenance-verifier/internal/platform"
)

func EncodeResult(v platform.Verification) ([]byte, error) {
	b, e := json.Marshal(v)
	if e != nil {
		return nil, fmt.Errorf("encode result: %w", e)
	}
	return b, nil
}
func DecodeResult(b []byte) (platform.Verification, error) {
	var v platform.Verification
	if e := json.Unmarshal(b, &v); e != nil {
		return v, fmt.Errorf("decode result: %w", e)
	}
	return v, nil
}
