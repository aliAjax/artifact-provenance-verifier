package adapter

import "context"

type Static struct{ Keys map[string][]byte }

func (s Static) Resolve(_ context.Context, id string) ([]byte, error) { return s.Keys[id], nil }
