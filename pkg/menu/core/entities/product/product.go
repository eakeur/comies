package product

import "gomies/pkg/sdk/types"

// Product is valuable resource to the client. It can be sold or used within the client's dependencies
type Product struct {
	types.Entity

	// Code represents how the store's crew call this product internally
	Code string

	// Name is the official name of the product, not exactly the name that the customer sees, but indeed the name
	// shown in fiscal documents
	Name string

	// CategoryID is the identifier of the category this product belongs to
	CategoryID types.ID

	// CategoryExternalID is the external identifier of the category this product belongs to
	CategoryExternalID types.External

	// CostPrice is how much the store pays to make or store this product
	CostPrice types.Currency

	// MaximumQuantity is how many unities of this product the store can support
	MaximumQuantity types.Quantity

	// MinimumQuantity is the lowest quantity of this product the store can have
	MinimumQuantity types.Quantity

	// Location is a brief description of where this stock is located
	Location string

	// Price is how much the customer pays for this product
	SalePrice types.Currency

	// Unit is the measure type that this product is sold
	SaleUnit types.UnitType

	// MinimumSale is the lowest number of unities of this product that can be sold
	MinimumSale types.Quantity

	// MaximumDiscount is how much discount a common operator can provide for the customer
	MaximumDiscount types.Percentage

	// Sale has useful information on how to sell this product
	Sale SaleInformation

	types.Store
}

func (p Product) Validate() error {

	if len(p.Code) < 2 || len(p.Code) > 12 {
		return ErrInvalidCode
	}

	if len(p.Name) < 2 || len(p.Name) > 60 {
		return ErrInvalidName
	}

	if p.SalePrice <= 0 {
		return ErrInvalidSalePrice
	}

	if p.MinimumSale <= types.QuantityMinimum {
		return ErrMinimumSaleQuantity
	}

	return nil
}
