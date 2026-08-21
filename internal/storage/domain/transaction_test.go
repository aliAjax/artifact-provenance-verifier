package domain

import (
	"context"
	"errors"
	"testing"
)

type failingTx struct{}

func (failingTx) Within(context.Context, func(context.Context) error) error {
	return errors.New("rollback failed")
}

type countingTx struct{ calls int }

func (t *countingTx) Within(ctx context.Context, fn func(context.Context) error) error {
	t.calls++
	return fn(ctx)
}

func TestUnitOfWorkPreservesOperationError(t *testing.T) {
	want := errors.New("operation failed")
	err := (UnitOfWork{Tx: failingTx{}}).Run(context.Background(), func(context.Context) error { return want })
	if !errors.Is(err, want) {
		t.Fatalf("error=%v", err)
	}
}

func TestUnitOfWorkRunsOperationOnce(t *testing.T) {
	tx := &countingTx{}
	calls := 0
	err := (UnitOfWork{Tx: tx}).Run(context.Background(), func(context.Context) error { calls++; return nil })
	if err != nil || calls != 1 || tx.calls != 1 {
		t.Fatalf("calls=%d txCalls=%d err=%v", calls, tx.calls, err)
	}
}

func TestUnitOfWorkRejectsNilCallback(t *testing.T) {
	if err := (UnitOfWork{}).Run(context.Background(), nil); err == nil {
		t.Fatal("nil transaction callback was accepted")
	}
}
