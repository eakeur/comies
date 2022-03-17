package crew

import "errors"

var (
	ErrWrongPassword = errors.New("the password provided is different from the stored one")
)
