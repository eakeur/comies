package menu

import (
	"comies/core/types"
	"comies/jobs/menu"
	"time"
)

type Ingredient struct {
	ProductID    types.ID       `json:"product_id"`
	IngredientID types.ID       `json:"ingredient_id"`
	Quantity     types.Quantity `json:"quantity"`
	Optional     bool           `json:"optional"`
}

type Movement struct {
	ID        types.ID       `json:"id"`
	ProductID types.ID       `json:"product_id"`
	AgentID   types.ID       `json:"agent_id"`
	Type      types.Type     `json:"type"`
	Date      time.Time      `json:"date"`
	Quantity  types.Quantity `json:"quantity"`
}

type StockBalance struct {
	// Balance is the amount stocked of this product
	Balance types.Quantity `json:"balance"`
}

type Price struct {
	Value types.Currency `json:"value"`
}

type ItemName struct {
	// Name is the official name of the product, not exactly the name that the customer sees, but indeed the name
	// shown in fiscal documents
	Name string `json:"name"`
}

type Item struct {
	// Code represents how the store's crew call this product internally
	Code string `json:"code"`
	// Name is the official name of the product, not exactly the name that the customer sees, but indeed the name
	// shown in fiscal documents
	Name string `json:"name"`
	// Type is the type of the product
	Type types.Type `json:"type"`
	// CostPrice is how much the store pays to make or store this product
	CostPrice types.Currency `json:"cost_price"`
	// Price is how much the customer pays for this product
	SalePrice types.Currency `json:"sale_price"`
	// Unit is the measure type that this product is sold
	SaleUnit types.UnitType `json:"sale_unit"`
	// MinimumSale is the lowest number of unities of this product that can be sold
	MinimumSale types.Quantity `json:"minimum_sale"`
	// MaximumQuantity is how many unities of this resource the stock can support
	MaximumQuantity types.Quantity `json:"maximum_quantity"`
	// MinimumQuantity is the lowest quantity of this resource the stock can have
	MinimumQuantity types.Quantity `json:"minimum_quantity"`
	// Location is a brief description of where this stock is located
	Location string `json:"location"`
}

type SaleableItem = menu.Saleable
