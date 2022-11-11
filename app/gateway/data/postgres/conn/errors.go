package conn

import "errors"

var (

	// ErrNoConnection throws when no transaction can be found in a given context
	ErrNoConnection = errors.New("could not find any transaction in this context")
)
