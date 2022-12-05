package product

import (
	"comies/core/menu/product"
)

var _ product.Actions = actions{}

type actions struct{}

func NewActions() product.Actions {
	return &actions{}
}
