package ingredient

import (
	"comies/core/menu/ingredient"
)

var _ ingredient.Actions = actions{}

type actions struct {
}

func NewActions() ingredient.Actions {
	return &actions{}
}
