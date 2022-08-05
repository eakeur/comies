package v1

import (
	"comies/app/core/entities/product"
	"comies/app/gateway/api/failures"
	"comies/app/gateway/api/handler"
	"comies/app/sdk/throw"
	"comies/app/sdk/types"
	"context"
	"encoding/json"
	"net/http"
)

// UpdateProduct updates a product to the store's menu.
//
// @Summary     Updates product
// @Description updates a product to the store's menu.
// @Tags        Product
// @Param       product body UpdateProductRequest true "The properties to define the product"
// @Success     204
// @Failure     400 {object} handler.Response{error=handler.Error{}} "INVALID_ID: Happens if the product id provided is not a valid one"
// @Failure     412 {object} handler.Response{error=handler.Error{}} "PRODUCT_CODE_ALREADY_EXISTS: Happens if the code provided is assigned to another product already"
// @Failure     422 {object} handler.Response{error=handler.Error{}} "PRODUCT_ZERO_SALE_QUANTITY: Happens if the minimum sale field is not greater than zero"
// @Failure     422 {object} handler.Response{error=handler.Error{}} "PRODUCT_ZERO_PRICE: Happens if the sale price field is not greater than zero"
// @Failure     422 {object} handler.Response{error=handler.Error{}} "PRODUCT_INVALID_CODE: Happens if the product code is not longer than 2 and shorter than 12 characters
// @Failure     422 {object} handler.Response{error=handler.Error{}} "PRODUCT_INVALID_NAME: Happens if the product name is not longer than 2 and shorter than 60 characters"
// @Failure     500 {object} handler.Response{error=handler.Error{}} "ERR_INTERNAL_SERVER_ERROR: Happens if an unexpected error happens on the API side"
// @Router      /menu/products/{product_id} [PUT]
func (s Service) UpdateProduct(ctx context.Context, r *http.Request) handler.Response {

	var p UpdateProductRequest
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		return handler.JSONParsingErrorResponse(err)
	}

	prod := p.ToProduct()

	id, err, res := handler.GetResourceIDFromURL(r, "product_id")
	if err != nil {
		return res
	}

	prod.ID = id

	err = s.menu.UpdateProduct(ctx, prod)
	if err != nil {
		return failures.Handle(throw.Error(err))
	}

	return handler.ResponseWithData(http.StatusNoContent, nil)
}

type UpdateProductRequest struct {
	// Code represents how the store's crew call this product internally
	Code string `json:"code"`
	// Name is the official name of the product, not exactly the name that the customer sees, but indeed the name
	// shown in fiscal documents
	Name string `json:"name"`
	// Type is the type of the product
	Type product.Type `json:"type"`
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

func (p UpdateProductRequest) ToProduct() product.Product {
	return product.Product{
		Code: p.Code,
		Name: p.Name,
		Type: p.Type,
		Sale: product.Sale{
			CostPrice:   p.CostPrice,
			SalePrice:   p.SalePrice,
			SaleUnit:    p.SaleUnit,
			MinimumSale: p.MinimumSale,
		},
		Stock: product.Stock{
			MaximumQuantity: p.MaximumQuantity,
			MinimumQuantity: p.MinimumQuantity,
			Location:        p.Location,
		},
	}
}