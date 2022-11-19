package price

import "comies/app/core/menu/price"

var _ price.Actions = actions{}

type actions struct{}

func NewActions() price.Actions {
	return actions{}
}
