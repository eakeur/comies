package order

import (
	"gomies/pkg/sdk/types"
	"time"
)

type (
	Status        int
	ItemStatus    int
	DeliverMethod int

	Order struct {
		types.Entity
		types.Store
		PlacedAt      time.Time
		AddressID     types.UID
		CustomerID    types.UID
		Status        Status
		DeliverMethod DeliverMethod
	}

	Item struct {
		types.Entity
		OrderID    types.ID
		OrderUID   types.UID
		ProductsID []types.UID
		Quantities []types.Quantity
		ItemStatus ItemStatus
		Price      types.Currency
		FinalPrice types.Currency
		Discount   types.Discount
	}
)

func (o Order) Validate() error {
	return nil
}

func (i Item) Validate() error {

	productsLen := len(i.ProductsID)
	quantitiesLen := len(i.Quantities)
	if productsLen == 0 {
		return ErrMissingProductIDs
	}

	if quantitiesLen == 0 {
		return ErrMissingQuantities
	}

	if quantitiesLen != productsLen {
		return ErrProductsAndQuantities
	}

	return nil
}
