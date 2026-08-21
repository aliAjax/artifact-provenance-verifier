package application

import (
	"context"
	"github.com/example/artifact-provenance-verifier/internal/webhook/domain"
	"testing"
)

func TestDrainReturnsIndependentEvents(t *testing.T) {
	o := NewOutbox()
	o.Add(domain.Event{Type: "first", ResourceID: "a"})
	got := o.Drain(context.Background())
	o.Add(domain.Event{Type: "second", ResourceID: "b"})
	if len(got) != 1 || got[0].Type != "first" {
		t.Fatalf("drained snapshot changed: %#v", got)
	}
	next := o.Drain(context.Background())
	if len(next) != 1 || next[0].Type != "second" {
		t.Fatalf("next batch wrong: %#v", next)
	}
}

func TestAddCopiesEventPayload(t *testing.T) {
	o := NewOutbox()
	payload := []byte("first")
	o.Add(domain.Event{Type: "payload", Payload: payload})
	payload[0] = 'X'
	got := o.Drain(context.Background())
	if string(got[0].Payload) != "first" {
		t.Fatalf("payload alias leaked: %q", got[0].Payload)
	}
}

func TestAddCopiesEventHeaders(t *testing.T) {
	o := NewOutbox()
	headers := map[string]string{"x-trace": "one"}
	o.Add(domain.Event{Type: "headers", Headers: headers})
	headers["x-trace"] = "two"
	got := o.Drain(context.Background())
	if got[0].Headers["x-trace"] != "one" {
		t.Fatalf("header alias leaked: %#v", got[0].Headers)
	}
}

func TestDrainCopiesNestedEventValues(t *testing.T) {
	o := NewOutbox()
	o.events = []domain.Event{{Type: "nested", Payload: []byte("stable"), Headers: map[string]string{"x": "stable"}}}
	original := o.events[0]
	got := o.Drain(context.Background())
	got[0].Payload[0] = 'X'
	got[0].Headers["x"] = "changed"
	if string(original.Payload) != "stable" || original.Headers["x"] != "stable" {
		t.Fatalf("drain returned nested aliases: %#v", original)
	}
}

func TestDrainReleasesQueueStorage(t *testing.T) {
	o := NewOutbox()
	o.Add(domain.Event{Type: "release"})
	_ = o.Drain(context.Background())
	if o.events != nil {
		t.Fatalf("drain retained queue backing storage: %#v", o.events)
	}
}
