package infrastructure

import "time"

type Expiry struct{ NotBefore, NotAfter time.Time }

func (e Expiry) Valid(now time.Time) bool {
	return (e.NotBefore.IsZero() || !now.Before(e.NotBefore)) && (e.NotAfter.IsZero() || now.Before(e.NotAfter))
}
