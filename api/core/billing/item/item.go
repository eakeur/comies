package item

import (
	"comies/core/types"
)

type Item struct {
	ID          types.ID       `json:"id"`
	BillID      types.ID       `json:"bill_id"`
	ReferenceID types.ID       `json:"reference_id"`
	Name        string         `json:"name"`
	Quantity    types.Quantity `json:"quantity"`
	UnitPrice   types.Currency `json:"unit_price"`
	Discounts   types.Currency `json:"discounts"`
}

func (i Item) Validate() (Item, error) {

	if err := i.BillID.Validate(); err != nil {
		return i, err
	}

	if err := i.ReferenceID.Validate(); err != nil {
		return i, err
	}

	if i.Quantity < 0 {
		return i, ErrInvalidQuantity
	}

	if i.UnitPrice < 0 {
		return i, ErrInvalidUnitPrice
	}
	return i, nil
}
