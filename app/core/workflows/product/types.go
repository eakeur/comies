package product

import (
	"gomies/app/core/types/currency"
	"gomies/app/core/types/id"
	"gomies/app/core/types/percentage"
	"gomies/app/core/types/quantity"
	"gomies/app/core/types/units"
	"time"
)

type CreateInput struct {
	Code                 string
	Name                 string
	CategoryID           id.External
	Display              string
	Description          string
	Price                currency.Currency
	Unit                 units.UnitType
	MinimumSale          quantity.Quantity
	MaximumDiscount      percentage.Percentage
	CostPrice            currency.Currency
	MaximumStockQuantity quantity.Quantity
	MinimumStockQuantity quantity.Quantity
	Location             string
}

type CreateOutput struct {
	ID        id.External
	CreatedAt time.Time
}
