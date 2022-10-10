package rest

import (
	"errors"
)

type ErrorBinding map[error]Response

var errUnexpected = errors.New("unexpected")

func (b ErrorBinding) Default(err Response) ErrorBinding {
	b[errUnexpected] = err
	return b
}

func Fail(err error) Response {
	for internal, external := range failures {
		if errors.Is(err, internal) {
			return external.Err(err)
		}
	}

	return failures[errUnexpected].Err(err)
}
