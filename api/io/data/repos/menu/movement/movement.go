package movement

import (
	"comies/core/menu/movement"
)

var _ movement.Actions = actions{}

type actions struct{}

func NewActions() movement.Actions {
	return actions{}
}
