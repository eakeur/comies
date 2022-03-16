package permission

import "errors"

var (
	ErrNotAllowed = errors.New("the actual session cannot operate the targeted function")
)
