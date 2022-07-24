package menu

import (
	"comies/app/core/entities/movement"
	"comies/app/core/entities/product"
	"comies/app/gateway/api/response"
	"comies/app/sdk/throw"
	"comies/app/sdk/types"
	"net/http"
	"strconv"
	"time"
)

type (
	Product struct {

		// ID is the unique identifier of this product
		ID string `json:"id,omitempty"`

		// Code represents how the store's crew call this product internally
		Code string `json:"code"`

		// Name is the official name of the product, not exactly the name that the customer sees, but indeed the name
		// shown in fiscal documents
		Name string `json:"name"`

		// Type is the type of the product
		Type product.Type `json:"type"`

		// CostPrice is how much the store pays to make or store this product
		CostPrice types.Currency `json:"cost_price,omitempty"`

		// Price is how much the customer pays for this product
		SalePrice types.Currency `json:"sale_price,omitempty"`

		// Unit is the measure type that this product is sold
		SaleUnit types.UnitType `json:"sale_unit,omitempty"`

		// MinimumSale is the lowest number of unities of this product that can be sold
		MinimumSale types.Quantity `json:"minimum_sale,omitempty"`

		// MaximumQuantity is how many unities of this resource the stock can support
		MaximumQuantity types.Quantity `json:"maximum_quantity,omitempty"`

		// MinimumQuantity is the lowest quantity of this resource the stock can have
		MinimumQuantity types.Quantity `json:"minimum_quantity,omitempty"`

		// Location is a brief description of where this stock is located
		Location string `json:"location,omitempty"`
	}

	Ingredient struct {
		// ID is the unique identifier of this ingredient
		ID           string         `json:"id,omitempty"`
		ProductID    string         `json:"product_id"`
		IngredientID string         `json:"ingredient_id"`
		Quantity     types.Quantity `json:"quantity"`
		Optional     bool           `json:"optional"`
	}

	Movement struct {
		// ID is the unique identifier of this movement
		ID string `json:"id,omitempty"`

		// ProductID is an identifier for the stock this movement references to
		ProductID string `json:"product_id,omitempty"`

		// Type points out if this movement is input or output
		Type movement.Type `json:"type,omitempty"`

		// Date is when the object got into the stock effectively
		Date time.Time `json:"date"`

		// Quantity is the amount being inserted or removed from this stock
		Quantity types.Quantity `json:"quantity,omitempty"`

		// PaidValue is how much was paid/received for this resource
		PaidValue types.Currency `json:"paid_value,omitempty"`

		// AgentID is the entity from this resource came from or is going to
		AgentID string `json:"agent_id,omitempty"`
	}

	AdditionResult struct {
		ID string `json:"id"`
	}

	MovementAdditionResult struct {
		ID      string         `json:"id"`
		Balance types.Quantity `json:"balance"`
	}

	ProductNameResult struct {
		Name string `json:"name"`
	}

	ProductStockBalanceResult struct {
		Balance types.Quantity `json:"balance"`
	}
)

func convertToID(in string) (types.ID, error, response.Response) {
	id, err := strconv.Atoi(in)
	if err != nil {
		return 0, err, response.WithError(http.StatusBadRequest, response.Error{
			Code: "INVALID_ID", Message: "The id provided is invalid",
		}).Err(throw.Error(err))
	}

	return types.ID(id), nil, response.Response{}
}
