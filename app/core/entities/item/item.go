package item

import (
	"gomies/app/sdk/fault"
	"gomies/app/sdk/types"
)

const (
	PreparingStatus Status = 0
	DoneStatus      Status = 10
	FailedStatus    Status = 20
)

type (
	Status int

	Item struct {
		ID         types.ID
		History    types.History
		Active     bool
		OrderID    types.ID
		Status     Status
		Price      types.Currency
		FinalPrice types.Currency
		Discount   types.Discount

		ProductID    types.ID
		Quantity     types.Quantity
		Observations string
		Details      Details
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
