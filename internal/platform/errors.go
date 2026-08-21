package platform

import "errors"

var (
	ErrNotFound     = errors.New("resource not found")
	ErrConflict     = errors.New("resource conflict")
	ErrInvalid      = errors.New("invalid request")
	ErrUnauthorized = errors.New("unauthorized")
	ErrDuplicate    = errors.New("duplicate submission")
)
