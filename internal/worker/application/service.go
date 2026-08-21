package application

import (
	"context"
	"time"
)

type Job interface{ Run(context.Context) error }
type Runner struct {
	Jobs     []Job
	Interval time.Duration
}

func (r *Runner) Run(ctx context.Context) {
	if r.Interval <= 0 {
		r.Interval = time.Minute
	}
	t := time.NewTicker(r.Interval)
	defer t.Stop()
	for {
		select {
		case <-ctx.Done():
			return
		case <-t.C:
			for _, j := range r.Jobs {
				_ = j.Run(ctx)
			}
		}
	}
}
