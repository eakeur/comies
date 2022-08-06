package menu

import (
	"comies/app/core/throw"
	"comies/app/gateway/api/handler"
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
// @Failure     404         {object} handler.Response{error=handler.Error{}} "PRODUCT_NOT_FOUND"
// @Failure     400         {object} handler.Response{error=handler.Error{}} "INVALID_ID"
// @Failure     500         {object} handler.Response{error=handler.Error{}} "ERR_INTERNAL_SERVER_ERROR"
// @Router      /menu/products/{product_id}/name [GET]
func (s Service) GetProductNameByID(ctx context.Context, r *http.Request) handler.Response {
	id, err := handler.GetResourceIDFromURL(r, "product_id")
	if err != nil {
		return handler.IDParsingErrorResponse(err)
	}

	name, err := s.menu.GetProductNameByID(ctx, id)
	if err != nil {
		return handler.Response{}
	}
	if err != nil {
		return handler.Fail(throw.Error(err))
	}

	return handler.ResponseWithData(http.StatusOK, GetProductNameResponse{Name: name})
}

type GetProductNameResponse struct {
	// Name is the official name of the product, not exactly the name that the customer sees, but indeed the name
	// shown in fiscal documents
	Name string `json:"name"`
}
