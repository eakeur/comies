package item

import (
	"gomies/app/sdk/fault"
	"gomies/app/sdk/types"
)

const (
	PreparingStatus Status = "PREPARING"
	DoneStatus      Status = "DONE"
	FailedStatus    Status = "FAILED"
)

type (
	Status string

	Item struct {
		ID           types.ID
		OrderID      types.ID
		Status       Status
		Price        types.Currency
		ProductID    types.ID
		Quantity     types.Quantity
		Observations string
		Details      Details
		History      types.History
	}

	Details struct {
		IgnoreIngredients  []Ignoring
		ReplaceIngredients []Replacement
	}

	Ignoring types.ID

	Replacement struct {
		From types.ID
		To   types.ID
	}
)

func (i Item) Validate() error {
	if i.Quantity <= types.QuantityMinimum {
		return fault.Wrap(ErrInvalidQuantity).
			DescribeF("the quantity should be bigger than %v", types.QuantityMinimum)
	}

	if i.ProductID.Empty() {
		return fault.Wrap(fault.ErrMissingID).Describe("a product id must be specified")
	}

	return nil
}
