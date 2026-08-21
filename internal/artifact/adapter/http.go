package adapter

import "net/http"

type Handler struct{ Next http.Handler }

func (h Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) { h.Next.ServeHTTP(w, r) }
