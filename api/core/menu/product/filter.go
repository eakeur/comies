package product

import (
	"comies/core/types"
)

type Filter struct {
	Types      []types.Type
	Code, Name string
}
