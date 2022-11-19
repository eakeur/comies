package prices

import (
	"comies/app/io/http/request"
	"comies/app/io/http/send"
	"context"
	"net/http"
)

// GetProductIngredients fetches all product ingredients.
//
// @Summary     Fetches ingredients
// @Description Fetches all product ingredients.
// @Tags        Product
// @Param       product_id path     string false "The product ID"
// @Success     200         {object} rest.Response{data=[]Ingredient{}}
// @Failure     500         {object} rest.Response{error=rest.Error{}} "ERR_INTERNAL_SERVER_ERROR"
// @Router      /menu/products/{product_id}/ingredients [GET]
func (h Handler) List(ctx context.Context, r request.Request) send.Response {
	productID, err := r.IDParam("product_id")
	if err != nil {
		return send.IDError(err)
	}

	list, err := h.prices.ListPrices(ctx, productID)
	if err != nil {
		return send.FromError(err)
	}

	return send.Data(http.StatusOK, list)
}
