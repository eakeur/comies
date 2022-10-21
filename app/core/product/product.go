package product

import (
	"comies/app/core/types"
)

type Product struct {
	ID              types.ID       `json:"id"`
	Code            string         `json:"code"`
	Name            string         `json:"name"`
	Type            types.Type     `json:"type"`
	Balance         types.Quantity `json:"balance"`
	CostPrice       types.Currency `json:"cost_price"`
	SalePrice       types.Currency `json:"sale_price"`
	SaleUnit        types.UnitType `json:"sale_unit"`
	MinimumSale     types.Quantity `json:"minimum_sale"`
	MaximumQuantity types.Quantity `json:"maximum_quantity"`
	MinimumQuantity types.Quantity `json:"minimum_quantity"`
	Location        string         `json:"location"`
}
