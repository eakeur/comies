package movements

import (
	"comies/api/request"
	"comies/api/send"
	"comies/core/menu/movement"
	"comies/core/types"
	"context"
	"net/http"
)

// GetProductStockBalance fetches a product name by its ID or code.
//
// @Summary     Fetches a product
// @Description Fetches a product name by its id.
// @Tags        Product
// @Param       product_id path     string false "The product ID"
// @Success     200         {object} rest.Response{data=GetProductBalanceResponse{}}
// @Failure     400         {object} rest.Response{error=rest.Error{}} "INVALID_ID"
// @Failure     500         {object} rest.Response{error=rest.Error{}} "ERR_INTERNAL_SERVER_ERROR"
// @Router      /menu/products/{product_id}/stock-balance [GET]
func (h Handler) GetProductStockBalance(ctx context.Context, r request.Request) send.Response {
	id, err := r.IDParam("product_id")
	if err != nil {
		return send.IDError(err)
	}

	bal, err := h.movements.GetProductStockBalance(ctx, movement.Filter{
		ProductID: id,
	})
	if err != nil {
		return send.FromError(err)
	}

	return send.Data(http.StatusOK, GetProductBalanceResponse{Balance: bal})
}

type GetProductBalanceResponse struct {
	// Balance is the amount stocked of this product
	Balance types.Quantity `json:"balance"`
}
