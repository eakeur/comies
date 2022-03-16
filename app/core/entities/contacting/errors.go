package contacting

import "errors"

var (
	ErrNotFound = errors.New("the address searched does not exist or could not be found")

	ErrAlreadyExists = errors.New("the address you are trying to create already exists")
)
