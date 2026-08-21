package api

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/example/artifact-provenance-verifier/internal/artifact/application"
	attapp "github.com/example/artifact-provenance-verifier/internal/attestation/application"
	"github.com/example/artifact-provenance-verifier/internal/platform"
	sbomapp "github.com/example/artifact-provenance-verifier/internal/sbom/application"
	"github.com/example/artifact-provenance-verifier/internal/storage/domain"
	verapp "github.com/example/artifact-provenance-verifier/internal/verification/application"
	vulnapp "github.com/example/artifact-provenance-verifier/internal/vulnerability/application"
	"log/slog"
	"net/http"
	"strconv"
	"strings"
	"sync/atomic"
	"time"
)

type Server struct {
	cfg       platform.Config
	artifacts *application.Service
	atts      *attapp.Service
	sboms     *sbomapp.Service
	vulns     *vulnapp.Service
	verify    *verapp.Service
	logger    *slog.Logger
	requests  atomic.Uint64
}

func New(cfg platform.Config, r domain.Repository, log *slog.Logger) *Server {
	return &Server{cfg: cfg, artifacts: application.New(r), atts: attapp.New(r), sboms: sbomapp.New(r), vulns: vulnapp.New(r), verify: verapp.New(r), logger: log}
}
func (s *Server) Handler() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/healthz", s.health)
	mux.HandleFunc("/readyz", s.health)
	mux.HandleFunc("/metrics", s.metrics)
	mux.HandleFunc("/v1/artifacts", s.artifactsHandler)
	mux.HandleFunc("/v1/artifacts/", s.artifactSub)
	mux.HandleFunc("/v1/vulnerabilities", s.vulnerabilityHandler)
	mux.HandleFunc("/v1/verifications", s.verificationHandler)
	mux.HandleFunc("/v1/recheck", s.verificationHandler)
	return s.middleware(mux)
}
func (s *Server) middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		s.requests.Add(1)
		if r.Method != "GET" && s.cfg.APIKey != "" && r.Header.Get("X-API-Key") != s.cfg.APIKey {
			writeErr(w, http.StatusUnauthorized, platform.ErrUnauthorized)
			return
		}
		r.Body = http.MaxBytesReader(w, r.Body, s.cfg.MaxBody)
		ctx, cancel := context.WithTimeout(r.Context(), s.cfg.RequestTimeout)
		defer cancel()
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
func (s *Server) health(w http.ResponseWriter, _ *http.Request) {
	writeJSON(w, 200, map[string]any{"status": "ok"})
}
func (s *Server) metrics(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	fmt.Fprintf(w, "provenance_http_requests_total %d\n", s.requests.Load())
}
func (s *Server) artifactsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		q := r.URL.Query().Get("q")
		a, e := s.artifacts.List(r.Context(), q)
		if e != nil {
			writeErr(w, 500, e)
			return
		}
		writeJSON(w, 200, a)
		return
	}
	if r.Method != http.MethodPost {
		writeErr(w, 405, fmt.Errorf("method not allowed"))
		return
	}
	var a platform.Artifact
	if !decode(w, r, &a) {
		return
	}
	x, e := s.artifacts.Register(r.Context(), a)
	if e != nil {
		writeErr(w, status(e), e)
		return
	}
	writeJSON(w, 201, x)
}
func (s *Server) artifactSub(w http.ResponseWriter, r *http.Request) {
	p := strings.TrimPrefix(r.URL.Path, "/v1/artifacts/")
	parts := strings.Split(strings.Trim(p, "/"), "/")
	if len(parts) == 0 || parts[0] == "" {
		writeErr(w, 404, platform.ErrNotFound)
		return
	}
	id := parts[0]
	if len(parts) == 1 && r.Method == http.MethodGet {
		a, e := s.artifacts.Get(r.Context(), id)
		if e != nil {
			writeErr(w, status(e), e)
			return
		}
		writeJSON(w, 200, a)
		return
	}
	if len(parts) < 2 {
		writeErr(w, 404, platform.ErrNotFound)
		return
	}
	switch parts[1] {
	case "attestations":
		s.attestationHandler(w, r, id)
	case "sboms":
		s.sbomHandler(w, r, id)
	case "impact":
		s.impactHandler(w, r, id)
	default:
		writeErr(w, 404, platform.ErrNotFound)
	}
}
func (s *Server) attestationHandler(w http.ResponseWriter, r *http.Request, id string) {
	if r.Method == http.MethodGet {
		a, e := s.atts.List(r.Context(), id)
		if e != nil {
			writeErr(w, 500, e)
			return
		}
		writeJSON(w, 200, a)
		return
	}
	var a platform.Attestation
	if !decode(w, r, &a) {
		return
	}
	a.ArtifactID = id
	x, e := s.atts.Submit(r.Context(), a)
	if e != nil {
		writeErr(w, status(e), e)
		return
	}
	writeJSON(w, 201, x)
}
func (s *Server) sbomHandler(w http.ResponseWriter, r *http.Request, id string) {
	if r.Method == http.MethodGet {
		x, e := s.sboms.Get(r.Context(), id)
		if e != nil {
			writeErr(w, status(e), e)
			return
		}
		writeJSON(w, 200, x)
		return
	}
	var x platform.SBOM
	if !decode(w, r, &x) {
		return
	}
	x.ArtifactID = id
	v, e := s.sboms.Submit(r.Context(), x)
	if e != nil {
		writeErr(w, status(e), e)
		return
	}
	writeJSON(w, 201, v)
}
func (s *Server) impactHandler(w http.ResponseWriter, r *http.Request, id string) {
	x, e := s.sboms.Get(r.Context(), id)
	if e != nil {
		writeErr(w, status(e), e)
		return
	}
	v, _ := s.vulns.Match(r.Context(), x)
	writeJSON(w, 200, map[string]any{"artifact_id": id, "vulnerabilities": v})
}
func (s *Server) vulnerabilityHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		v, e := s.vulns.List(r.Context())
		if e != nil {
			writeErr(w, 500, e)
			return
		}
		writeJSON(w, 200, v)
		return
	}
	var v platform.Vulnerability
	if !decode(w, r, &v) {
		return
	}
	x, e := s.vulns.Add(r.Context(), v)
	if e != nil {
		writeErr(w, status(e), e)
		return
	}
	writeJSON(w, 201, x)
}
func (s *Server) verificationHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		writeErr(w, 405, fmt.Errorf("method not allowed"))
		return
	}
	var in verapp.Inputs
	if !decode(w, r, &in) {
		return
	}
	v, e := s.verify.Verify(r.Context(), in)
	if e != nil {
		writeErr(w, status(e), e)
		return
	}
	writeJSON(w, 200, v)
}
func decode(w http.ResponseWriter, r *http.Request, v any) bool {
	if e := json.NewDecoder(r.Body).Decode(v); e != nil {
		writeErr(w, 400, e)
		return false
	}
	return true
}
func writeJSON(w http.ResponseWriter, code int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	_ = json.NewEncoder(w).Encode(v)
}
func writeErr(w http.ResponseWriter, code int, e error) {
	writeJSON(w, code, map[string]any{"error": e.Error()})
}
func status(e error) int {
	if errors.Is(e, platform.ErrNotFound) {
		return 404
	}
	if errors.Is(e, platform.ErrConflict) || errors.Is(e, platform.ErrDuplicate) {
		return 409
	}
	if errors.Is(e, platform.ErrUnauthorized) {
		return 401
	}
	if strings.Contains(strings.ToLower(e.Error()), "required") || strings.Contains(strings.ToLower(e.Error()), "invalid") {
		return 400
	}
	return 422
}
func ParsePort(addr string) int {
	p := strings.LastIndex(addr, ":")
	if p < 0 {
		return 0
	}
	n, _ := strconv.Atoi(addr[p+1:])
	return n
}
func Shutdown(ctx context.Context, srv *http.Server) error {
	c, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	return srv.Shutdown(c)
}
