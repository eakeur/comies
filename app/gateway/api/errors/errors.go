package errors

import (
	"errors"
)

type ErrorBinding map[error]error

var errUnexpected = errors.New("unexpected")

func (b ErrorBinding) HandleError(err error) error {
	for internal, external := range b {
		if errors.Is(err, internal) {
			return external
		}
	}

	return b[errUnexpected]
}

func (b ErrorBinding) Default(err error) ErrorBinding {
	b[errUnexpected] = err

	return b
}
