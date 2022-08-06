package menu

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

// CreateProduct adds a product to the store's menu.
//
// @Summary     Create product
// @Description Adds a product to the store's menu.
// @Tags        Product
// @Param       product body     CreateProductRequest true "The properties to define the product"
// @Success     201     {object} handler.Response{data=CreateProductResponse{}}
// @Failure     412     {object} handler.Response{error=handler.Error{}} "PRODUCT_CODE_ALREADY_EXISTS"
// @Failure     422     {object} handler.Response{error=handler.Error{}} "PRODUCT_ZERO_SALE_QUANTITY, PRODUCT_ZERO_PRICE, PRODUCT_INVALID_CODE, PRODUCT_INVALID_NAME, PRODUCT_INVALID_TYPE"
// @Failure     500     {object} handler.Response{error=handler.Error{}} "ERR_INTERNAL_SERVER_ERROR"
// @Router      /menu/products [POST]
func (s Service) CreateProduct(ctx context.Context, r *http.Request) handler.Response {

	var p CreateProductRequest
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		return handler.JSONParsingErrorResponse(err)
	}

	prd, err := s.menu.CreateProduct(ctx, p.ToProduct())
	if err != nil {
		return failures.Handle(throw.Error(err))
	}

	return handler.ResponseWithData(http.StatusCreated, CreateProductResponse{ID: prd.ID.String()})
}

type (
	CreateProductRequest struct {
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

	CreateProductResponse struct {
		ID string `json:"id"`
	}
)

func (p CreateProductRequest) ToProduct() product.Product {
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
