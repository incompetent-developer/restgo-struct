package auth

import (
	"errors"
)

// Error variables
var (
	ErrTokenExpired = errors.New("Token expired")
	ErrTokenInvalid = errors.New("Token invalid")
)
