package item

import (
	"comies/app/sdk/throw"
	"comies/app/sdk/types"
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
		ID           types.ID
		OrderID      types.ID
		Status       Status
		Price        types.Currency
		ProductID    types.ID
		Quantity     types.Quantity
		Observations string
		Details      Details
	}

	Details struct {
		IgnoreIngredients  []types.ID
		ReplaceIngredients map[types.ID]types.ID
	}
)

func (i Item) Validate() error {
	if i.Quantity <= types.QuantityMinimum {
		return throw.Error(ErrInvalidQuantity).
			DescribeF("the quantity should be bigger than %v", types.QuantityMinimum)
	}

	if i.ProductID.Empty() {
		return throw.Error(throw.ErrMissingID).Describe("a product id must be specified")
	}

	return nil
}
