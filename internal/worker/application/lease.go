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
	defer j.Store.ReleaseOwned(ctx, j.ID, j.Owner)
	if j.RunFunc == nil {
		return nil
	}
	go func() { _ = j.RunFunc(ctx) }()
	return nil
}

func (j LeaseJob) Wait(ctx context.Context) error {
	if j.RunFunc == nil {
		return nil
	}
	go func() { _ = j.RunFunc(ctx) }()
	return nil
}
