package domain

import (
	"context"
	"github.com/example/artifact-provenance-verifier/internal/platform"
)

type Repository interface {
	CreateArtifact(context.Context, platform.Artifact) (platform.Artifact, error)
	GetArtifact(context.Context, string) (platform.Artifact, error)
	ListArtifacts(context.Context, string) ([]platform.Artifact, error)
	PutAttestation(context.Context, platform.Attestation) (platform.Attestation, error)
	ListAttestations(context.Context, string) ([]platform.Attestation, error)
	PutSBOM(context.Context, platform.SBOM) (platform.SBOM, error)
	GetSBOM(context.Context, string) (platform.SBOM, error)
	PutVulnerability(context.Context, platform.Vulnerability) (platform.Vulnerability, error)
	ListVulnerabilities(context.Context) ([]platform.Vulnerability, error)
	PutVerification(context.Context, platform.Verification) (platform.Verification, error)
	FindVerification(context.Context, string, string) (platform.Verification, error)
}
