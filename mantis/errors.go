package mantis

import "errors"

var (
	ErrUnauthorized      = errors.New("unauthorized")
	ErrInvalidParameters = errors.New("invalid parameters")
	ErrServerError       = errors.New("server error")
	ErrNetworkError      = errors.New("network error")
)
