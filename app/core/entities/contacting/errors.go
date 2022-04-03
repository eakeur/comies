package contacting

import "errors"

var (
	ErrMissingResourceID = errors.New("the resource id informed is empty or invalid")
)
