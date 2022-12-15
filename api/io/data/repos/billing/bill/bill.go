package bill

import "comies/core/billing/bill"

var _ bill.Actions = actions{}

type actions struct{}

func NewActions() bill.Actions {
	return actions{}
}
