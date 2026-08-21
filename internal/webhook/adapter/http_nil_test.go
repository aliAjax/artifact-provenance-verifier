package adapter

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHTTPSendWithNilClient(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) { w.WriteHeader(http.StatusNoContent) }))
	defer srv.Close()
	if err := (&HTTP{}).Send(context.Background(), srv.URL, nil); err != nil {
		t.Fatalf("nil client send failed: %v", err)
	}
}

func TestHTTPSendWithNilReceiver(t *testing.T) {
	var h *HTTP
	if err := h.Send(context.Background(), "https://example.invalid", nil); err == nil {
		t.Fatal("nil HTTP receiver did not report an error")
	}
}
