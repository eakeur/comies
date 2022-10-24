package product

import (
	"comies/app/core/types"
)

type Filter struct {
	Type       types.Type
	Code, Name string
}
