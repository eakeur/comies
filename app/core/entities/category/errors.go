package category

import "errors"

var (
	ErrInvalidName = errors.New("the name is not valid. please provide one longer than 2 (two) and shorter than 60 (sixty) characters")

)
