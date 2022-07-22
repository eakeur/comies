package response

import "errors"

type ErrorBinding map[error]Response

var errUnexpected = errors.New("unexpected")

func (b ErrorBinding) Handle(err error) Response {
	for internal, external := range b {
		if errors.Is(err, internal) {
			return external.Err(err)
		}
	}

	return b[errUnexpected].Err(err)
}

func (b ErrorBinding) Default(err Response) ErrorBinding {
	b[errUnexpected] = err
	return b
}
