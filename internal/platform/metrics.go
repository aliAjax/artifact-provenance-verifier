package platform

import (
	"fmt"
	"sync/atomic"
)

type Metrics struct {
	Requests      atomic.Uint64
	Errors        atomic.Uint64
	Verifications atomic.Uint64
	BytesIn       atomic.Uint64
	BytesOut      atomic.Uint64
}

func (m *Metrics) Text() string {
	return fmt.Sprintf("provenance_requests_total %d\nprovenance_errors_total %d\nprovenance_verifications_total %d\nprovenance_bytes_in_total %d\nprovenance_bytes_out_total %d\n", m.Requests.Load(), m.Errors.Load(), m.Verifications.Load(), m.BytesIn.Load(), m.BytesOut.Load())
}
