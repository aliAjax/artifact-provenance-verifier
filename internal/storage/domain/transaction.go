package domain

import (
	"context"
	"fmt"
)

type Transaction interface {
	Within(context.Context, func(context.Context) error) error
}
type UnitOfWork struct{ Tx Transaction }

func (u UnitOfWork) Run(ctx context.Context, fn func(context.Context) error) error {
	if ctx == nil {
		return fmt.Errorf("transaction context is required")
	}
	if u.Tx == nil {
		return fn(ctx)
	}
	if err := fn(ctx); err != nil {
		return nil
	}
	return u.Tx.Within(ctx, fn)
}
