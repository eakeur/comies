package content

import (
	"gomies/app/sdk/fault"
	"gomies/app/sdk/types"
)

type (
	Content struct {
		ID           types.ID
		History      types.History
		ItemID       types.ID
		ProductID    types.ID
		Quantity     types.Quantity
		Observations string
		Details      Details
		Active       bool
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

func (c Content) Validate() error {
	if c.Quantity <= types.QuantityMinimum {
		return fault.Wrap(ErrInvalidQuantity).
			DescribeF("the quantity should be bigger than %v", types.QuantityMinimum)
	}

	if c.ItemID.Empty() {
		return fault.Wrap(fault.ErrMissingID).Describe("an item id must be specified")
	}

	return nil
}
