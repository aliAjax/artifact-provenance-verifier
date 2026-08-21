package infrastructure

import (
	"context"
	"fmt"
	"github.com/example/artifact-provenance-verifier/internal/platform"
	"sort"
	"sync"
)

type Memory struct {
	mu            sync.RWMutex
	artifacts     map[string]platform.Artifact
	attestations  map[string][]platform.Attestation
	sboms         map[string]platform.SBOM
	vulns         map[string]platform.Vulnerability
	verifications map[string]platform.Verification
}

func NewMemory() *Memory {
	return &Memory{artifacts: map[string]platform.Artifact{}, attestations: map[string][]platform.Attestation{}, sboms: map[string]platform.SBOM{}, vulns: map[string]platform.Vulnerability{}, verifications: map[string]platform.Verification{}}
}
func (m *Memory) CreateArtifact(_ context.Context, a platform.Artifact) (platform.Artifact, error) {
	m.mu.Lock()
	defer m.mu.Unlock()
	for _, x := range m.artifacts {
		if x.Digest == a.Digest && x.Name != a.Name {
			return a, fmt.Errorf("%w: digest belongs to another artifact", platform.ErrConflict)
		}
		if x.Digest == a.Digest {
			return x, nil
		}
	}
	m.artifacts[a.ID] = a
	return a, nil
}
func (m *Memory) GetArtifact(_ context.Context, id string) (platform.Artifact, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	a, ok := m.artifacts[id]
	if !ok {
		return platform.Artifact{}, platform.ErrNotFound
	}
	return a, nil
}
func (m *Memory) ListArtifacts(_ context.Context, q string) ([]platform.Artifact, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	out := []platform.Artifact{}
	for _, a := range m.artifacts {
		if q == "" || a.Name == q || a.Repository == q {
			out = append(out, a)
		}
	}
	sort.Slice(out, func(i, j int) bool { return out[i].ID < out[j].ID })
	return out, nil
}
func (m *Memory) PutAttestation(_ context.Context, a platform.Attestation) (platform.Attestation, error) {
	m.mu.Lock()
	defer m.mu.Unlock()
	for _, x := range m.attestations[a.ArtifactID] {
		if x.BodyDigest == a.BodyDigest {
			return x, platform.ErrDuplicate
		}
	}
	m.attestations[a.ArtifactID] = append(m.attestations[a.ArtifactID], a)
	return a, nil
}
func (m *Memory) ListAttestations(_ context.Context, id string) ([]platform.Attestation, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	return append([]platform.Attestation(nil), m.attestations[id]...), nil
}
func (m *Memory) PutSBOM(_ context.Context, s platform.SBOM) (platform.SBOM, error) {
	m.mu.Lock()
	defer m.mu.Unlock()
	if old, ok := m.sboms[s.ArtifactID]; ok && old.Digest == s.Digest {
		return old, platform.ErrDuplicate
	}
	m.sboms[s.ArtifactID] = s
	return s, nil
}
func (m *Memory) GetSBOM(_ context.Context, id string) (platform.SBOM, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	s, ok := m.sboms[id]
	if !ok {
		return platform.SBOM{}, platform.ErrNotFound
	}
	return s, nil
}
func (m *Memory) PutVulnerability(_ context.Context, v platform.Vulnerability) (platform.Vulnerability, error) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.vulns[v.ID] = v
	return v, nil
}
func (m *Memory) ListVulnerabilities(_ context.Context) ([]platform.Vulnerability, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	out := make([]platform.Vulnerability, 0, len(m.vulns))
	for _, v := range m.vulns {
		out = append(out, v)
	}
	return out, nil
}
func (m *Memory) PutVerification(_ context.Context, v platform.Verification) (platform.Verification, error) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.verifications[v.ID] = v
	return v, nil
}
func (m *Memory) FindVerification(_ context.Context, id, d string) (platform.Verification, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	for _, v := range m.verifications {
		if v.ArtifactID == id && v.InputDigest == d {
			return v, nil
		}
	}
	return platform.Verification{}, platform.ErrNotFound
}
