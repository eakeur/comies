package item

import (
	"comies/core/types"
)

type Item struct {
	ID          types.ID       `json:"id"`
	BillID      types.ID       `json:"bill_id"`
	ReferenceID types.ID       `json:"reference_id"`
	Debts       types.Currency `json:"debts"`
	Credits     types.Currency `json:"credits"`
	Description types.Text     `json:"description"`
}

func (c Item) WithID(id types.ID) Item {
	c.ID = id
	return c
}

func (c Item) WithDescription(description types.Text) Item {
	c.Description = description
	return c
}

func (c Item) WithCredits(v types.Currency) Item {
	c.Credits = v
	if v < 0 {
		c.Credits = v * -1
	}

	return c
}

func (c Item) WithDebts(v types.Currency) Item {
	c.Debts = v
	if v > 0 {
		c.Debts = v * -1
	}

	return c
}

func (c Item) Validate() (Item, error) {
	if c.Credits < 0 || c.Debts > 0 || (c.Credits == 0 && c.Debts == 0) {
		return c, ErrInvalidCreditsOrDebts
	}
	return c, nil
}
