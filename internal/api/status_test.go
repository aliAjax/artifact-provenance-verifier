package api

import (
	"errors"
	"fmt"
	"testing"

	"github.com/example/artifact-provenance-verifier/internal/platform"
)

func TestNotFoundErrorStatusPreserved(t *testing.T) {
	wrapped := fmt.Errorf("lookup artifact: %w", platform.ErrNotFound)
	if got := status(wrapped); got != 404 {
		t.Fatalf("status=%d, want 404", got)
	}
	if !errors.Is(wrapped, platform.ErrNotFound) {
		t.Fatal("test setup lost sentinel")
	}
}

func TestConflictErrorStatusPreserved(t *testing.T) {
	if got := status(fmt.Errorf("write: %w", platform.ErrConflict)); got != 409 {
		t.Fatalf("status=%d", got)
	}
}

func TestDuplicateErrorStatusPreserved(t *testing.T) {
	if got := status(fmt.Errorf("submit: %w", platform.ErrDuplicate)); got != 409 {
		t.Fatalf("status=%d", got)
	}
}

func TestUnauthorizedErrorStatusPreserved(t *testing.T) {
	if got := status(fmt.Errorf("auth: %w", platform.ErrUnauthorized)); got != 401 {
		t.Fatalf("status=%d", got)
	}
}
