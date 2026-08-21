package adapter

import "context"

type HealthChecker interface{ Ping(context.Context) error }
type SQLStore struct{ DSN string }

func NewSQLStore(dsn string) *SQLStore         { return &SQLStore{DSN: dsn} }
func (s *SQLStore) Ping(context.Context) error { return nil }
