package product

import (
	"gomies/app/core/types/currency"
	"gomies/app/core/types/percentage"
	"gomies/app/core/types/quantity"
	"gomies/app/core/types/units"
)

// SaleInformation wraps rules for selling this product
type SaleInformation struct {

	// Display is what will be shown for customers as the name of the product
	Display string

	// Description is a text that describes the product for the customer
	Description string

	// Price is how much the customer pays for this product
	Price currency.Currency

	// Unit is the measure type that this product is sold
	Unit units.UnitType

	// MinimumSale is the lowest number of unities of this product that can be sold
	MinimumSale quantity.Quantity

	// MaximumDiscount is how much discount a common operator can provide for the customer
	MaximumDiscount percentage.Percentage
}

func (sd SaleInformation) Validate() error {

	if sd.Price <= 0 {
		return ErrInvalidSalePrice
	}

	if sd.MinimumSale <= quantity.Minimum {
		return ErrMinimumSaleQuantity
	}

	return nil
}
