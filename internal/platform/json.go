package platform

import (
	"bytes"
	"encoding/json"
	"fmt"
)

func MarshalCanonical(v any) ([]byte, error) {
	b, e := json.Marshal(v)
	if e != nil {
		return nil, e
	}
	var buf bytes.Buffer
	if e = json.Compact(&buf, b); e != nil {
		return nil, fmt.Errorf("compact json: %w", e)
	}
	return buf.Bytes(), nil
}
func DecodeStrict(data []byte, v any) error {
	d := json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if e := d.Decode(v); e != nil {
		return fmt.Errorf("decode json: %w", e)
	}
	return nil
}
func IsJSON(data []byte) bool { var v any; return json.Unmarshal(data, &v) == nil }
func MergeMaps(a, b map[string]any) map[string]any {
	out := map[string]any{}
	for k, v := range a {
		out[k] = v
	}
	for k, v := range b {
		out[k] = v
	}
	return out
}
