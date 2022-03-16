package crew

import "errors"

var (
	ErrNotFound = errors.New("the operator searched does not exist or could not be found")

	ErrAlreadyExists = errors.New("the operator you are trying to create already exists")

	ErrHasDependants = errors.New("the operator being deleted has things depending on them")

	ErrWrongPassword = errors.New("the password provided is different from the stored one")
)
