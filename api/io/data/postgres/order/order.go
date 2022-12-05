package order

import (
	"comies/core/ordering/order"
)

var _ order.Actions = actions{}

type actions struct{}

func NewActions() order.Actions {
	return actions{}
}
