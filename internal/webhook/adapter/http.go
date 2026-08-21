package adapter

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	"time"
)

type HTTP struct{ Client *http.Client }

func New() *HTTP { return &HTTP{Client: &http.Client{Timeout: 5 * time.Second}} }
func (h *HTTP) Send(ctx context.Context, url string, p []byte) error {
	req, e := http.NewRequestWithContext(ctx, http.MethodPost, url, bytes.NewReader(p))
	if e != nil {
		return e
	}
	req.Header.Set("Content-Type", "application/json")
	resp, e := h.Client.Do(req)
	if e != nil {
		return e
	}
	defer resp.Body.Close()
	io.Copy(io.Discard, resp.Body)
	if resp.StatusCode >= 300 {
		return fmt.Errorf("webhook status %d", resp.StatusCode)
	}
	return nil
}
