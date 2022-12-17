package item

import (
	"comies/core/billing/item"
)

var _ item.Actions = actions{}

type actions struct{}

func NewActions() item.Actions {
	return actions{}
}
