package application

import (
	"context"
	"testing"
)

type nilSender struct{}

func (*nilSender) Send(context.Context, string, []byte) error { return nil }

type derefSender struct{ calls int }

func (s *derefSender) Send(context.Context, string, []byte) error {
	s.calls++
	return nil
}

func TestNotifyTypedNilSenderPanicsBeforeValidation(t *testing.T) {
	var sender *derefSender
	if err := New(sender).Notify(context.Background(), "https://example.invalid", nil); err == nil {
		t.Fatal("typed nil sender was accepted")
	}
}

func TestNotifyTypedNilSender(t *testing.T) {
	var sender *nilSender
	s := New(sender)
	defer func() {
		if r := recover(); r != nil {
			t.Fatalf("typed nil sender panicked: %v", r)
		}
	}()
	if err := s.Notify(context.Background(), "https://example.invalid", nil); err == nil {
		t.Fatal("typed nil sender was accepted")
	}
}

func TestNotifyNilServiceReceiver(t *testing.T) {
	var s *Service
	if err := s.Notify(context.Background(), "https://example.invalid", nil); err == nil {
		t.Fatal("nil service receiver was accepted")
	}
}

func TestNotifyNilSenderReportsConfigurationError(t *testing.T) {
	if err := New(nil).Notify(context.Background(), "https://example.invalid", nil); err == nil {
		t.Fatal("nil sender did not return an error")
	}
}
