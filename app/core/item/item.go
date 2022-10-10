package item

import (
	"comies/app/core/id"
	"comies/app/core/types"
)

const (
	NoStatus        Status = 0
	PreparingStatus Status = 10
	DoneStatus      Status = 20
	FailedStatus    Status = 30
)

type (
	Status int

	Item struct {
		ID           id.ID
		OrderID      id.ID
		Status       Status
		Price        types.Currency
		ProductID    id.ID
		Quantity     types.Quantity
		Observations string
		Details      Details
	}

	Details struct {
		IgnoreIngredients  []id.ID
		ReplaceIngredients map[id.ID]id.ID
	}
)

func (s Status) Validate() error {
	switch s {
	case PreparingStatus, DoneStatus, FailedStatus:
		return nil
	default:
		return ErrInvalidStatus
	}
}

func (i Item) Validate() error {
	if err := i.Status.Validate(); err != nil {
		return err
	}

	if i.Quantity <= types.QuantityMinimum {
		return ErrInvalidQuantity
	}

	if i.ProductID.Empty() {
		return types.ErrMissingID
	}

	return nil
}
