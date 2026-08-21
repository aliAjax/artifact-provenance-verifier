package application

import (
	"context"
	"github.com/example/artifact-provenance-verifier/internal/worker/adapter"
)

type LeaseJob struct {
	Store     *adapter.LeaseStore
	ID, Owner string
	RunFunc   func(context.Context) error
}

func (j LeaseJob) Run(ctx context.Context) error {
	if !j.Store.Acquire(ctx, j.ID, j.Owner) {
		return nil
	}
	defer j.Store.Release(ctx, j.ID)
	if j.RunFunc == nil {
		return nil
	}
	return j.RunFunc(ctx)
}
