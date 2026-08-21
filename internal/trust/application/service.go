package application

import "context"

type KeyResolver interface {
	Resolve(context.Context, string) ([]byte, error)
}
type Service struct{ Resolver KeyResolver }

func New(r KeyResolver) *Service { return &Service{Resolver: r} }
