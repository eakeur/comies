package product

import (
	"comies/app/core/menu/product"
)

var _ product.Actions = actions{}

type actions struct{}

func NewActions() product.Actions {
	return &actions{}
}
