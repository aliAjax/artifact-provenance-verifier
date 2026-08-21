package platform

import "testing"

func TestTransitionRejectsIllegalRestore(t *testing.T) {
	for _, p := range [][2]Status{{StatusWithdrawn, StatusActive}, {StatusWithdrawn, StatusIsolated}} {
		if err := Transition(p[0], p[1]); err == nil {
			t.Fatalf("allowed %s -> %s", p[0], p[1])
		}
	}
}

func TestTransitionAllowsWithdrawnRestore(t *testing.T) {
	if err := Transition(StatusWithdrawn, StatusRestored); err != nil {
		t.Fatalf("restore rejected: %v", err)
	}
}

func TestTransitionRejectsRestoredWithdrawn(t *testing.T) {
	if err := Transition(StatusRestored, StatusWithdrawn); err == nil {
		t.Fatal("restored artifact was withdrawn directly")
	}
}

func TestTransitionRejectsRestoredActive(t *testing.T) {
	if err := Transition(StatusRestored, StatusActive); err == nil {
		t.Fatal("restored artifact moved backwards")
	}
}

func TestTransitionRejectsActiveRestore(t *testing.T) {
	if err := Transition(StatusActive, StatusRestored); err == nil {
		t.Fatal("active artifact was restored directly")
	}
}

func TestTransitionRejectsIsolatedRestore(t *testing.T) {
	if err := Transition(StatusIsolated, StatusRestored); err == nil {
		t.Fatal("isolated artifact was restored directly")
	}
}

func TestTransitionRejectsActiveWithdrawn(t *testing.T) {
	if err := Transition(StatusActive, StatusWithdrawn); err == nil {
		t.Fatal("active artifact was withdrawn without isolation")
	}
}

func TestTransitionRejectsInvalidSource(t *testing.T) {
	if err := Transition(Status("unknown"), StatusActive); err == nil {
		t.Fatal("invalid source was accepted")
	}
}
