package failures

import (
	"comies/app/gateway/api/handler"
	"errors"
)

type ErrorBinding map[error]handler.Response

var errUnexpected = errors.New("unexpected")

func (b ErrorBinding) Default(err handler.Response) ErrorBinding {
	b[errUnexpected] = err
	return b
}

func Handle(err error) handler.Response {
	for internal, external := range failures {
		if errors.Is(err, internal) {
			return external.Err(err)
		}
	}

	return failures[errUnexpected].Err(err)
}
