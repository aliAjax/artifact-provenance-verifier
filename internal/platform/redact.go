package platform

import (
	"log/slog"
	"strings"
)

type Redacted string

func (r Redacted) LogValue() slog.Value { return slog.StringValue("[REDACTED]") }
func Redact(v string) string {
	if len(v) <= 8 {
		return "[REDACTED]"
	}
	return v[:4] + "..." + v[len(v)-4:]
}
func LooksSensitive(k string) bool {
	l := strings.ToLower(k)
	return strings.Contains(l, "token") || strings.Contains(l, "secret") || strings.Contains(l, "password") || strings.Contains(l, "private") || strings.Contains(l, "signature")
}
