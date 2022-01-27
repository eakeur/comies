package product

import (
	"gomies/app/core/types/currency"
	"gomies/app/core/types/quantity"
	"time"
)

// StockInformation wraps information about this product stock
type StockInformation struct {

	// CostPrice is how much the store pays to make or store this product
	CostPrice currency.Currency

	// LastUpdate is the date the stock was last updated
	LastUpdate time.Time

	// MaximumQuantity is how many unities of this product the store can support
	MaximumQuantity quantity.Quantity

	// MinimumQuantity is the lowest quantity of this product the store can have
	MinimumQuantity quantity.Quantity

	// Location is a brief description of where this stock is located
	Location string
}

func (sd StockInformation) Validate() error {
	return nil
}
