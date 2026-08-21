package application

import (
	"context"
	"errors"
	"sync"
	"testing"
	"time"

	"github.com/example/artifact-provenance-verifier/internal/worker/adapter"
)

func exerciseLeaseReaders(store *adapter.LeaseStore, id string) {
	start := make(chan struct{})
	var wg sync.WaitGroup
	for i := 0; i < 2; i++ {
		wg.Add(1)
		go func() { defer wg.Done(); <-start; _ = store.Owner(id) }()
	}
	close(start)
	wg.Wait()
}

func TestLeaseJobWaitsForCompletion(t *testing.T) {
	ctx := context.Background()
	started, release, finished := make(chan struct{}), make(chan struct{}), make(chan struct{})
	j := LeaseJob{Store: adapter.New(), ID: "job-1", Owner: "worker", RunFunc: func(got context.Context) error {
		close(started)
		<-release
		close(finished)
		return nil
	}}
	exerciseLeaseReaders(j.Store, "job-1")
	result := make(chan error, 1)
	go func() { result <- j.Run(ctx) }()
	<-started
	select {
	case <-result:
		t.Fatal("lease job returned before work completed")
	case <-time.After(50 * time.Millisecond):
	}
	close(release)
	if err := <-result; err != nil {
		t.Fatal(err)
	}
	select {
	case <-finished:
	case <-time.After(time.Second):
		t.Fatal("work did not finish")
	}
}

func TestLeaseJobPropagatesCallbackError(t *testing.T) {
	want := errors.New("callback failed")
	store := adapter.New()
	exerciseLeaseReaders(store, "other")
	j := LeaseJob{Store: store, ID: "job-error", Owner: "worker", RunFunc: func(context.Context) error { return want }}
	if err := j.Run(context.Background()); !errors.Is(err, want) {
		t.Fatalf("Run error=%v", err)
	}
}

func TestLeaseAcquireRejectsCanceledContext(t *testing.T) {
	store := adapter.New()
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	exerciseLeaseReaders(store, "job-cancel")
	_ = store.Acquire(ctx, "job-cancel", "worker")
	if store.Owner("job-cancel") != "" {
		t.Fatalf("canceled acquire stored owner=%q", store.Owner("job-cancel"))
	}
}

func TestLeaseReleaseRequiresOwner(t *testing.T) {
	store := adapter.New()
	_ = store.Acquire(context.Background(), "job-owned", "owner")
	exerciseLeaseReaders(store, "job-owned")
	_ = store.ReleaseOwned(context.Background(), "job-owned", "intruder")
	if store.Owner("job-owned") != "owner" {
		t.Fatalf("intruder released lease, owner=%q", store.Owner("job-owned"))
	}
}

func TestLeaseRenewRequiresOwner(t *testing.T) {
	store := adapter.New()
	_ = store.Acquire(context.Background(), "job-renew", "owner")
	exerciseLeaseReaders(store, "job-renew")
	_ = store.Renew(context.Background(), "job-renew", "intruder")
	if store.Owner("job-renew") != "owner" {
		t.Fatalf("intruder renewed lease, owner=%q", store.Owner("job-renew"))
	}
}
