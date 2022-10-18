package menu

import (
	"comies/app/api/request"
	"comies/app/api/send"
	"comies/app/core/types"
	"comies/app/jobs/menu"
	"context"
)

// CreateProduct adds a product to the store's menu.
//
// @Summary     Create product
// @Description Adds a product to the store's menu.
// @Tags        Product
// @Param       product body     CreateProductRequest true "The properties to define the product"
// @Success     201     {object} rest.Response{data=CreateProductResponse{}}
// @Failure     412     {object} rest.Response{error=rest.Error{}} "PRODUCT_CODE_ALREADY_EXISTS"
// @Failure     422     {object} rest.Response{error=rest.Error{}} "PRODUCT_ZERO_SALE_QUANTITY, PRODUCT_ZERO_PRICE, PRODUCT_INVALID_CODE, PRODUCT_INVALID_NAME, PRODUCT_INVALID_TYPE"
// @Failure     500     {object} rest.Response{error=rest.Error{}} "ERR_INTERNAL_SERVER_ERROR"
// @Router      /menu/products [POST]
func CreateProduct(ctx context.Context, r request.Request) send.Response {

	var p CreateProductRequest
	err := r.JSONBody(&p)
	if err != nil {
		return send.JSONError(err)
	}

	prd, err := menu.SaveProduct(ctx, menu.Product{
		Code:            p.Code,
		Name:            p.Name,
		Type:            p.Type,
		CostPrice:       p.CostPrice,
		SalePrice:       p.SalePrice,
		SaleUnit:        p.SaleUnit,
		MinimumSale:     p.MinimumSale,
		MaximumQuantity: p.MaximumQuantity,
		MinimumQuantity: p.MinimumQuantity,
		Location:        p.Location,
	})
	if err != nil {
		return send.FromError(err)
	}

	return send.CreatedWithID(prd.ID)
}

type CreateProductRequest struct {
	// Code represents how the store's crew call this product internally
	Code string `json:"code"`
	// Name is the official name of the product, not exactly the name that the customer sees, but indeed the name
	// shown in fiscal documents
	Name string `json:"name"`
	// Type is the type of the product
	Type menu.Type `json:"type"`
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
