package id

import "errors"

var (
	ErrNoID     = errors.New("this resource's id was not provided or is invalid")
)