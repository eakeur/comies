package v1

import (
	"comies/app/core/entities/movement"
	"comies/app/gateway/api/failures"
	"comies/app/gateway/api/handler"
	"comies/app/sdk/throw"
	"comies/app/sdk/types"
	"context"
	"net/http"
)

// GetProductStockBalance fetches a product name by its ID or code.
//
// @Summary     Fetches a product
// @Description Fetches a product name by its id.
// @Tags        Product
// @Param       product_key path     string false "The product ID"
// @Success     200         {object} handler.Response{data=GetProductBalanceResponse{}}
// @Failure     400         {object} handler.Response{error=handler.Error{}} "INVALID_ID: Happens if the product id provided is not a valid one"
// @Failure     500         {object} handler.Response{error=handler.Error{}} "ERR_INTERNAL_SERVER_ERROR: Happens if an unexpected error happens on the API side"
// @Router      /menu/products/{product_id}/stock-balance [GET]
func (s Service) GetProductStockBalance(ctx context.Context, r *http.Request) handler.Response {
	id, err, res := handler.GetResourceIDFromURL(r, "product_id")
	if err != nil {
		return res
	}

	bal, err := s.menu.GetMovementsBalance(ctx, movement.Filter{
		ProductID: id,
	})
	if err != nil {
		return failures.Handle(throw.Error(err))
	}

	return handler.ResponseWithData(http.StatusOK, GetProductBalanceResponse{Balance: bal})
}

type GetProductBalanceResponse struct {
	// Balance is the amount stocked of this product
	Balance types.Quantity `json:"balance"`
}
