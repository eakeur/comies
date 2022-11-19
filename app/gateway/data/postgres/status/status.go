package status

import "comies/app/core/ordering/status"

var _ status.Actions = actions{}

type actions struct{}

func NewActions() status.Actions {
	return actions{}
}
