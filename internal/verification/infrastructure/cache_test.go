package infrastructure

import (
	"github.com/example/artifact-provenance-verifier/internal/platform"
	"testing"
)

func TestCacheReturnsIndependentVerification(t *testing.T) {
	c := NewCache()
	v := platform.Verification{Checks: []platform.CheckResult{{Name: "initial"}}}
	c.Put("k", v)
	got, _ := c.Get("k")
	got.Checks[0].Name = "caller"
	again, _ := c.Get("k")
	if len(again.Checks) != 1 || again.Checks[0].Name != "initial" {
		t.Fatalf("cache was changed: %#v", again.Checks)
	}
}

func TestCacheCopiesInputChecks(t *testing.T) {
	c := NewCache()
	v := platform.Verification{Checks: []platform.CheckResult{{Name: "initial"}}}
	c.Put("k", v)
	v.Checks[0].Name = "caller"
	got, _ := c.Get("k")
	if got.Checks[0].Name != "initial" { t.Fatalf("cache input alias leaked: %#v", got.Checks) }
}

func TestCacheReturnsIndependentDetails(t *testing.T) {
	c := NewCache()
	c.Put("k", platform.Verification{Checks: []platform.CheckResult{{Name: "check", Details: map[string]any{"owner": "cache"}}}})
	got, _ := c.Get("k")
	got.Checks[0].Details["owner"] = "caller"
	again, _ := c.Get("k")
	if again.Checks[0].Details["owner"] != "cache" { t.Fatalf("details alias leaked: %#v", again.Checks[0].Details) }
}
