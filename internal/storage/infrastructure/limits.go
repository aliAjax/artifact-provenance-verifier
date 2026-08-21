package infrastructure

import "io"

type LimitedReader struct {
	R io.Reader
	N int64
}

func (l *LimitedReader) Read(p []byte) (int, error) {
	if l.N <= 0 {
		return 0, io.EOF
	}
	if int64(len(p)) > l.N {
		p = p[:l.N]
	}
	n, e := l.R.Read(p)
	l.N -= int64(n)
	return n, e
}
