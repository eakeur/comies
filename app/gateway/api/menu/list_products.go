package menu

import (
	"comies/app/core/entities/product"
	"comies/app/gateway/api/handler"
	"comies/app/sdk/throw"
	"comies/app/sdk/types"
	"context"
	"net/http"
	"strconv"
)

// ListProducts fetches a product by its ID or code.
//
// @Summary     Fetches a product
// @Description Fetches a product by one of itd unique keys (id or code).
// @Tags        Product
// @Param       code query    string false "Adds a filter looking for the products codes"
// @Param       name query    string false "Adds a filter looking for the products names"
// @Param       type query    int    false "Adds a filter looking for the products types"
// @Param       out  query    bool   false "Searches products running out of stock only"
// @Success     200  {object} handler.Response{data=[]ListProductsResponse{}}
// @Success     200  {object} handler.Response{data=[]ListRunningOutProductsResponse{}}
// @Failure     500  {object} handler.Response{error=handler.Error{}} "ERR_INTERNAL_SERVER_ERROR"
// @Router      /menu/products [GET]
func (s Service) ListProducts(ctx context.Context, r *http.Request) handler.Response {
	query := r.URL.Query()
	filter := product.Filter{
		Code: query.Get("code"),
		Name: query.Get("name"),
		Type: 0,
	}

	runningOut := query.Get("out") == "true"

	t, err := strconv.Atoi(query.Get("type"))
	if err == nil {
		filter.Type = product.Type(t)
	}

	if runningOut {
		products, err := s.menu.ListProductsRunningOut(ctx)
		if err != nil {
			return handler.Fail(throw.Error(err))
		}

		return handler.ResponseWithData(http.StatusOK, NewListRunningOutProductsResponse(products))
	}

	products, err := s.menu.ListProducts(ctx, filter)
	if err != nil {
		return handler.Fail(throw.Error(err))
	}

	return handler.ResponseWithData(http.StatusOK, NewListProductsResponse(products))
}

type (
	ListProductsResponse struct {
		// ID is the unique identifier of this product
		ID string `json:"id"`
		// Code represents how the store's crew call this product internally
		Code string `json:"code"`
		// Name is the official name of the product, not exactly the name that the customer sees, but indeed the name
		// shown in fiscal documents
		Name string `json:"name"`
		// Type is the type of the product
		Type product.Type `json:"type"`
	}

	ListRunningOutProductsResponse struct {
		// ID is the unique identifier of this product
		ID string `json:"id"`
		// Code represents how the store's crew call this product internally
		Code string `json:"code"`
		// Name is the official name of the product, not exactly the name that the customer sees, but indeed the name
		// shown in fiscal documents
		Name string `json:"name"`
		// Type is the type of the product
		Type product.Type `json:"type"`
		// Unit is the measure type that this product is sold
		SaleUnit types.UnitType `json:"sale_unit"`
		// MaximumQuantity is how many unities of this resource the stock can support
		MaximumQuantity types.Quantity `json:"maximum_quantity"`
		// MinimumQuantity is the lowest quantity of this resource the stock can have
		MinimumQuantity types.Quantity `json:"minimum_quantity"`
		// Balance is the stock balance of this product
		Balance types.Quantity `json:"balance"`
	}
)

func NewListProductsResponse(list []product.Product) []ListProductsResponse {
	products := make([]ListProductsResponse, len(list))
	for i, p := range list {
		products[i] = ListProductsResponse{
			ID:   p.ID.String(),
			Code: p.Code,
			Name: p.Name,
			Type: p.Type,
		}
	}
	return products
}

func NewListRunningOutProductsResponse(list []product.Product) []ListRunningOutProductsResponse {
	products := make([]ListRunningOutProductsResponse, len(list))
	for i, p := range list {
		products[i] = ListRunningOutProductsResponse{
			ID:              p.ID.String(),
			Code:            p.Code,
			Name:            p.Name,
			Type:            p.Type,
			SaleUnit:        p.SaleUnit,
			MaximumQuantity: p.MaximumQuantity,
			MinimumQuantity: p.MinimumQuantity,
			Balance:         p.Balance,
		}
	}
	return products
}
