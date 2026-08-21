package platform

import (
	"fmt"
	"net/url"
	"regexp"
	"strings"
	"time"
)

var idPattern = regexp.MustCompile(`^[a-zA-Z0-9][a-zA-Z0-9._:/-]{0,255}$`)

func ValidateID(v string) error {
	if !idPattern.MatchString(v) {
		return fmt.Errorf("invalid identifier")
	}
	return nil
}
func ValidateURL(raw string) error {
	u, e := url.Parse(raw)
	if e != nil {
		return fmt.Errorf("url parse: %w", e)
	}
	if u.Scheme != "https" && u.Scheme != "http" {
		return fmt.Errorf("unsupported url scheme")
	}
	if u.Host == "" {
		return fmt.Errorf("url host required")
	}
	return nil
}
func NormalizeName(v string) string { return strings.ToLower(strings.TrimSpace(v)) }
func ValidateTimeWindow(start, end time.Time) error {
	if start.IsZero() || end.IsZero() {
		return nil
	}
	if end.Before(start) {
		return fmt.Errorf("end before start")
	}
	if end.Sub(start) > 365*24*time.Hour {
		return fmt.Errorf("time window too large")
	}
	return nil
}
func ValidateTags(tags []string) error {
	seen := map[string]bool{}
	for _, t := range tags {
		t = strings.TrimSpace(t)
		if t == "" || len(t) > 64 {
			return fmt.Errorf("invalid tag")
		}
		if seen[t] {
			return fmt.Errorf("duplicate tag")
		}
		seen[t] = true
	}
	return nil
}
func ContainsAny(v string, parts []string) bool {
	for _, p := range parts {
		if p != "" && strings.Contains(v, p) {
			return true
		}
	}
	return false
}
func SafeMessage(v string) string {
	v = strings.ReplaceAll(v, "\n", " ")
	if len(v) > 512 {
		return v[:512]
	}
	return v
}
