package product

import (
	"comies/app/core/types"
)

type Filter struct {
	Types      []types.Type
	Code, Name string
}
