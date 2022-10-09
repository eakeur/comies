package throw

import (
	"errors"
)

var (
	ErrNotFound      = errors.New("this resource could not be found or does not exist")
	ErrMissingID     = errors.New("this resource's id was not provided or is invalid")
	ErrAlreadyExists = errors.New("this resource key is already assigned to another resource")
)
