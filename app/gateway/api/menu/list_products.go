package menu

import (
	"comies/app/core/entities/product"
	"comies/app/core/throw"
	"comies/app/gateway/api/handler"
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
// @Success     200  {object} handler.Response{data=[]GetProductByKeyResponse{}}
// @Failure     500  {object} handler.Response{error=handler.Error{}} "ERR_INTERNAL_SERVER_ERROR"
// @Router      /menu/products [GET]
func (s Service) ListProducts(ctx context.Context, r *http.Request) handler.Response {
	query := r.URL.Query()
	filter := product.Filter{
		Code: query.Get("code"),
		Name: query.Get("name"),
		Type: 0,
	}

	t, err := strconv.Atoi(query.Get("type"))
	if err == nil {
		filter.Type = product.Type(t)
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
)

func NewListProductsResponse(list []product.Product) []GetProductByKeyResponse {
	products := make([]GetProductByKeyResponse, len(list))
	for i, p := range list {
		products[i] = NewGetProductByKeyResponse(p)
	}
	return products
}
