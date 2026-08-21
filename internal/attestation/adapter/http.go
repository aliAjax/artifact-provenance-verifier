package adapter

import "net/http"

func DecodeLimit(r *http.Request, n int64) { r.Body = http.MaxBytesReader(nil, r.Body, n) }
