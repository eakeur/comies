package conn

import "errors"

var (

	// ErrNoConnection throws when no transaction can be found in a given context
	ErrNoConnection = errors.New("could not find any transaction in this context")
)

type (
	Occurence interface {
		Check(err error) error
	}

	ConstraintError struct {
		Name, Code string
		Error      error
	}

	NoRowsError struct {
		Error error
	}
)

func MapDatabaseError(err error, occurrences ...Occurence) error {
	if len(occurrences) == 0 {
		return err
	}

	if e := occurrences[0].Check(err); e != nil {
		return e
	}

	return MapDatabaseError(err, occurrences[:1]...)
}
