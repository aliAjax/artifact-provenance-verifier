package api

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type ErrorBody struct {
	Code      string `json:"code"`
	Message   string `json:"message"`
	RequestID string `json:"request_id,omitempty"`
}

func encode(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(v)
}
func decodeStrict(r *http.Request, v any) error {
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if e := d.Decode(v); e != nil {
		return fmt.Errorf("invalid JSON: %w", e)
	}
	return nil
}
func method(w http.ResponseWriter, r *http.Request, want string) bool {
	if r.Method != want {
		w.Header().Set("Allow", want)
		encode(w, 405, ErrorBody{Code: "method_not_allowed", Message: "method not allowed"})
		return false
	}
	return true
}
func acceptJSON(r *http.Request) bool { return r.Header.Get("Content-Type") == "application/json" }
