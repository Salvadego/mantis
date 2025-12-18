package mantis

import (
	"errors"
	"fmt"
	"strings"
)

var (
	ErrUnauthorized      = errors.New("unauthorized")
	ErrInvalidParameters = errors.New("invalid parameters")
	ErrServerError       = errors.New("server error")
	ErrNetworkError      = errors.New("network error")
)

type APIError struct {
	StatusCode int
	Response   ErrorsResponse
}

func (e *APIError) Error() string {
	if !e.Response.IsError() || len(e.Response.Errors) == 0 {
		return fmt.Sprintf("API error: HTTP %d", e.StatusCode)
	}

	main := e.Response.Errors[0].Message
	var detailMsgs []string
	for _, d := range e.Response.Errors[0].Details {
		if d.Message != "" {
			detailMsgs = append(detailMsgs, d.Message)
		}
	}

	if len(detailMsgs) > 0 {
		return fmt.Sprintf("%s %s", main, strings.Join(detailMsgs, " "))
	}

	return main
}
