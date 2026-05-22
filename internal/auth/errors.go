package auth

import "errors"

var (
	ErrUnauthorized     = errors.New("unauthorized")
	ErrForbidden        = errors.New("forbidden")
	ErrSessionNotFound  = errors.New("session not found")
	ErrNoPrincipalFound = errors.New("no principal found")
)
