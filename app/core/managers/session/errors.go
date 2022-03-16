package session

import "errors"

var (
	ErrNoSession = errors.New("there is no access found in the context informed")

	ErrSessionInvalidOrExpired = errors.New("the token informed is invalid or has expired")
)
