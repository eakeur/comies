package products

import (
	"comies/app/api/request"
	"comies/app/api/send"
	"context"
	"net/http"
)

// GetProductNameByID fetches a product name by its ID or code.
//
// @Summary     Fetches a product
// @Description Fetches a product name by its id.
// @Tags        Product
// @Param       product_key path     string false "The product ID"
// @Success     200         {object} rest.Response{data=GetProductNameResponse{}}
// @Failure     404         {object} rest.Response{error=rest.Error{}} "PRODUCT_NOT_FOUND"
// @Failure     400         {object} rest.Response{error=rest.Error{}} "INVALID_ID"
// @Failure     500         {object} rest.Response{error=rest.Error{}} "ERR_INTERNAL_SERVER_ERROR"
// @Router      /menu/products/{product_id}/name [GET]
func (h Handler) GetNameByID(ctx context.Context, r request.Request) send.Response {
	id, err := r.IDParam("product_id")
	if err != nil {
		return send.IDError(err)
	}

	name, err := h.products.GetProductNameByID(ctx, id)
	if err != nil {
		return send.FromError(err)
	}

	return send.Data(http.StatusOK, GetProductNameResponse{Name: name})
}

type GetProductNameResponse struct {
	// Name is the official name of the product, not exactly the name that the customer sees, but indeed the name
	// shown in fiscal documents
	Name string `json:"name"`
}
