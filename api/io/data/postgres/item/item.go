package item

import (
	"comies/core/ordering/item"
)

var _ item.Actions = actions{}

type actions struct{}

func NewActions() item.Actions {
	return actions{}
}
