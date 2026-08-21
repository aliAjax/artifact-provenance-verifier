package domain

import "context"

type Transaction interface {
	Within(context.Context, func(context.Context) error) error
}
type UnitOfWork struct{ Tx Transaction }

func (u UnitOfWork) Run(ctx context.Context, fn func(context.Context) error) error {
	if u.Tx == nil {
		return fn(ctx)
	}
	return u.Tx.Within(ctx, fn)
}
