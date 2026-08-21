package application

import (
	"context"
	"fmt"
	"github.com/example/artifact-provenance-verifier/internal/platform"
	"github.com/example/artifact-provenance-verifier/internal/storage/domain"
	"time"
)

type Inputs struct {
	ArtifactID         string   `json:"artifact_id"`
	RequireAttestation bool     `json:"require_attestation"`
	RequireSBOM        bool     `json:"require_sbom"`
	MaxSeverity        string   `json:"max_severity"`
	AllowedSources     []string `json:"allowed_sources"`
	AllowedLicenses    []string `json:"allowed_licenses"`
}
type Service struct{ Repo domain.Repository }

func New(r domain.Repository) *Service { return &Service{Repo: r} }
func (s *Service) Verify(ctx context.Context, in Inputs) (platform.Verification, error) {
	a, e := s.Repo.GetArtifact(ctx, in.ArtifactID)
	if e != nil {
		return platform.Verification{}, e
	}
	atts, _ := s.Repo.ListAttestations(ctx, a.ID)
	sbom, sbomErr := s.Repo.GetSBOM(ctx, a.ID)
	checks := []platform.CheckResult{}
	if in.RequireAttestation {
		status := "pass"
		msg := "valid attestation found"
		if len(atts) == 0 {
			status = "fail"
			msg = "no attestation submitted"
		}
		checks = append(checks, platform.CheckResult{Name: "attestation", Severity: "high", Status: status, Message: msg})
	} else {
		checks = append(checks, platform.CheckResult{Name: "attestation", Severity: "low", Status: "skip", Message: "optional"})
	}
	if in.RequireSBOM {
		status := "pass"
		msg := "SBOM available"
		if sbomErr != nil || len(sbom.Components) == 0 {
			status = "fail"
			msg = "SBOM missing or empty"
		}
		checks = append(checks, platform.CheckResult{Name: "sbom", Severity: "high", Status: status, Message: msg})
	}
	if sbomErr == nil && in.MaxSeverity != "" {
		vulns, _ := s.Repo.ListVulnerabilities(ctx)
		bad := 0
		for _, v := range vulns {
			if v.Withdrawn {
				continue
			}
			for _, c := range sbom.Components {
				if c.Name == v.PackagePattern && platform.SeverityRank(v.Severity) >= platform.SeverityRank(in.MaxSeverity) {
					bad++
				}
			}
		}
		status := "pass"
		msg := "no blocking vulnerabilities"
		if bad > 0 {
			status = "fail"
			msg = fmt.Sprintf("%d blocking vulnerabilities", bad)
		}
		checks = append(checks, platform.CheckResult{Name: "vulnerabilities", Severity: in.MaxSeverity, Status: status, Message: msg, Details: map[string]any{"count": bad}})
	}
	conclusion := "pass"
	for _, c := range checks {
		if c.Status == "fail" {
			conclusion = "fail"
		}
	}
	digest, _ := platform.HashJSON(struct {
		A platform.Artifact
		I Inputs
		S platform.SBOM
	}{a, in, sbom})
	if old, err := s.Repo.FindVerification(ctx, a.ID, digest); err == nil {
		old.Reused = true
		return old, nil
	}
	v := platform.Verification{ID: fmt.Sprintf("ver-%d", time.Now().UnixNano()), ArtifactID: a.ID, InputDigest: digest, Conclusion: conclusion, RuleVersion: "v1", Checks: checks, CreatedAt: time.Now().UTC()}
	return s.Repo.PutVerification(ctx, v)
}
func (s *Service) Get(ctx context.Context, id, d string) (platform.Verification, error) {
	return s.Repo.FindVerification(ctx, id, d)
}
