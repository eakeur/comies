package product

import "comies/app/core/types"

type (
	Product struct {
		ID types.ID
		// Code represents how the store's crew call this product internally
		Code string

		// Name is the official name of the product, not exactly the name that the customer sees, but indeed the name
		// shown in fiscal documents
		Name string

		Type Type

		Balance types.Quantity

		Sale

		Stock
	}

	Sale struct {
		// CostPrice is how much the store pays to make or store this product
		CostPrice types.Currency

		// Price is how much the customer pays for this product
		SalePrice types.Currency

		// Unit is the measure type that this product is sold
		SaleUnit types.UnitType

		// MinimumSale is the lowest number of unities of this product that can be sold
		MinimumSale types.Quantity
	}

	Stock struct {
		// MaximumQuantity is how many unities of this resource the stock can support
		MaximumQuantity types.Quantity
		// MinimumQuantity is the lowest quantity of this resource the stock can have
		MinimumQuantity types.Quantity
		// Location is a brief description of where this stock is located
		Location string
	}

	Type int
)

const (
	NoType              Type = 0
	OutputType          Type = 10
	OutputCompositeType Type = 20
	InputType           Type = 30
	InputCompositeType  Type = 40
)

func (t Type) Validate() error {
	switch t {
	case InputType, OutputType, InputCompositeType, OutputCompositeType:
		return nil
	default:
		return ErrInvalidType
	}
}

func (p Product) Validate() error {

	if len(p.Code) < 2 || len(p.Code) > 12 {
		return ErrInvalidCode
	}

	if len(p.Name) < 2 || len(p.Name) > 60 {
		return ErrInvalidName
	}

	if err := p.Type.Validate(); err != nil {
		return err
	}

	if (p.SalePrice <= 0 && p.Type == OutputType) || p.CostPrice <= 0 {
		return ErrInvalidPrice
	}

	if p.Type == OutputType && p.MinimumSale <= types.QuantityMinimum {
		return ErrMinimumSaleQuantity
	}

	return nil
}
