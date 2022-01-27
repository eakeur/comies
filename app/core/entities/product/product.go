package product

import (
	"gomies/app/core/types/entity"
	"gomies/app/core/types/id"
)

// Product is valuable resource to the client. It can be sold or used within the client's dependencies
type Product struct {
	entity.Entity

	// Code represents how the store's crew call this product internally
	Code string

	// Name is the official name of the product, not exactly the name that the customer sees, but indeed the name
	// shown in fiscal documents
	Name string

	// CategoryID is the identifier of the category this product belongs to
	CategoryID id.ID

	// CategoryExternalID is the external identifier of the category this product belongs to
	CategoryExternalID id.External

	// Stock has useful information on how to store this product
	Stock StockInformation

	// Sale has useful information on how to sell this product
	Sale SaleInformation
}

func (p Product) Validate() error {

	if len(p.Code) < 2 || len(p.Code) > 12 {
		return ErrInvalidCode
	}

	if len(p.Name) < 2 || len(p.Name) > 60 {
		return ErrInvalidName
	}

	if err := p.Stock.Validate(); err != nil {
		return err
	}

	if err := p.Sale.Validate(); err != nil {
		return err
	}

	return nil
}
