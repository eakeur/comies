package product

import "gomies/pkg/sdk/types"

type (
	Product struct {
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

		Type Type

		Stock

		Sale

		types.Store
	}

	Stock struct {
		// CostPrice is how much the store pays to make or store this product
		CostPrice types.Currency

		// MaximumQuantity is how many unities of this product the store can support
		MaximumQuantity types.Quantity

		// MinimumQuantity is the lowest quantity of this product the store can have
		MinimumQuantity types.Quantity

		// Location is a brief description of where this stock is located
		Location string
	}

	Sale struct {
		// Price is how much the customer pays for this product
		SalePrice types.Currency

		// Unit is the measure type that this product is sold
		SaleUnit types.UnitType

		// MinimumSale is the lowest number of unities of this product that can be sold
		MinimumSale types.Quantity

		// MaximumDiscount is how much discount a common operator can provide for the customer
		MaximumDiscount types.Percentage

		// Display is what will be shown for customers as the name of the product
		Display string

		// Description is a text that describes the product for the customer
		Description string

		// HasIngredients tells if the product has other products chained to it
		HasIngredients bool
	}

	Type int
)

const (
	OutputType Type = iota
	InputType  Type = iota
)

func (p Product) Validate() error {

	if len(p.Code) < 2 || len(p.Code) > 12 {
		return ErrInvalidCode
	}

	if len(p.Name) < 2 || len(p.Name) > 60 {
		return ErrInvalidName
	}

	if (p.SalePrice <= 0 && p.Type == OutputType) || p.CostPrice <= 0 {
		return ErrInvalidPrice
	}

	if p.Type == OutputType && p.MinimumSale <= types.QuantityMinimum {
		return ErrMinimumSaleQuantity
	}

	return nil
}
