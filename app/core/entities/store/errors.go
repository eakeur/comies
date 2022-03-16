package store

import "errors"

var (
	ErrNotFound = errors.New("the store searched does not exist or could not be found")

	ErrAlreadyExists = errors.New("the store you are trying to create already exists")

	ErrHasDependants = errors.New("the store being deleted has things depending on it")
)
