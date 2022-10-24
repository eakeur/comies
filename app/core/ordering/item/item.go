package item

import (
	"comies/app/core/types"
)

type Item struct {
	ID           types.ID       `json:"id"`
	OrderID      types.ID       `json:"order_id"`
	Status       types.Status   `json:"status"`
	ProductID    types.ID       `json:"product_id"`
	Quantity     types.Quantity `json:"quantity"`
	Observations string         `json:"observations"`
}

func (i Item) WithID(id types.ID) Item {
	i.ID = id
	return i
}

func (i Item) Validate() (Item, error) {
	if err := ValidateItemStatus(i.Status); err != nil {
		return i, err
	}

	if i.Quantity <= types.QuantityMinimum {
		return i, ErrInvalidQuantity
	}

	if err := i.ProductID.Validate(); err != nil {
		return i, err
	}

	if err := i.OrderID.Validate(); err != nil {
		return i, err
	}
	return i, nil
}
