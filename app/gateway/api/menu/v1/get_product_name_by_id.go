package v1

import (
	"comies/app/gateway/api/failures"
	"comies/app/gateway/api/handler"
	"comies/app/sdk/throw"
	"context"
	"net/http"
)

// GetProductNameByID fetches a product name by its ID or code.
//
// @Summary     Fetches a product
// @Description Fetches a product name by its id.
// @Tags        Product
// @Param       product_key path     string false "The product ID"
// @Success     200         {object} handler.Response{data=GetProductNameResponse{}}
// @Failure     404         {object} handler.Response{error=handler.Error{}} "PRODUCT_NOT_FOUND: Happens if the product could not be found or does not exist"
// @Failure     400         {object} handler.Response{error=handler.Error{}} "INVALID_ID: Happens if the product id provided is not a valid one"
// @Failure     500         {object} handler.Response{error=handler.Error{}} "ERR_INTERNAL_SERVER_ERROR: Happens if an unexpected error happens on the API side"
// @Router      /menu/products/{product_id}/name [GET]
func (s Service) GetProductNameByID(ctx context.Context, r *http.Request) handler.Response {
	id, err, res := handler.GetResourceIDFromURL(r, "product_id")
	if err != nil {
		return res
	}

	name, err := s.menu.GetProductNameByID(ctx, id)
	if err != nil {
		return handler.Response{}
	}
	if err != nil {
		return failures.Handle(throw.Error(err))
	}

	return handler.ResponseWithData(http.StatusOK, GetProductNameResponse{Name: name})
}

type GetProductNameResponse struct {
	// Name is the official name of the product, not exactly the name that the customer sees, but indeed the name
	// shown in fiscal documents
	Name string `json:"name"`
}
