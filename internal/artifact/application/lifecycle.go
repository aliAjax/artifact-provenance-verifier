package application

import (
	"context"
	"fmt"
	"github.com/example/artifact-provenance-verifier/internal/platform"
	"github.com/example/artifact-provenance-verifier/internal/storage/domain"
)

type Lifecycle struct{ Repo domain.Repository }

func NewLifecycle(r domain.Repository) *Lifecycle { return &Lifecycle{Repo: r} }
func (l *Lifecycle) SetStatus(ctx context.Context, id, status string) (platform.Artifact, error) {
	a, e := l.Repo.GetArtifact(ctx, id)
	if e != nil {
		return a, e
	}
	switch status {
	case "active", "isolated", "withdrawn", "restored":
	default:
		return a, fmt.Errorf("unsupported status")
	}
	if a.Status == "withdrawn" && status != "restored" {
		return a, fmt.Errorf("withdrawn artifact must be restored first")
	}
	a.Status = status
	return l.Repo.CreateArtifact(ctx, a)
}
func (l *Lifecycle) AddTag(ctx context.Context, id, tag string) (platform.Artifact, error) {
	a, e := l.Repo.GetArtifact(ctx, id)
	if e != nil {
		return a, e
	}
	if e = platform.ValidateTags(append(a.Tags, tag)); e != nil {
		return a, e
	}
	a.Tags = platform.UniqueStrings(append(a.Tags, tag))
	return l.Repo.CreateArtifact(ctx, a)
}
