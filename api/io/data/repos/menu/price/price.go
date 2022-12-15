package price

import "comies/core/menu/price"

var _ price.Actions = actions{}

type actions struct{}

func NewActions() price.Actions {
	return actions{}
}
