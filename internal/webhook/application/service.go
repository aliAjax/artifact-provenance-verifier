package application

import (
	"context"
)

type Sender interface {
	Send(context.Context, string, []byte) error
}
type Service struct{ Sender Sender }

func New(s Sender) *Service { return &Service{Sender: s} }
func (s *Service) Notify(ctx context.Context, url string, payload []byte) error {
	if s.Sender == nil {
		return nil
	}
	return s.Sender.Send(ctx, url, payload)
}
