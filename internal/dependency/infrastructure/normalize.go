package infrastructure

import "strings"

func NormalizePURL(v string) string { return strings.ToLower(strings.TrimSpace(v)) }
func PackageName(v string) string {
	v = NormalizePURL(v)
	if i := strings.Index(v, "@"); i >= 0 {
		v = v[:i]
	}
	return strings.TrimPrefix(v, "pkg:")
}
