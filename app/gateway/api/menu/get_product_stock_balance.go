package menu

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
// @Param       product_id path     string false "The product ID"
// @Success     200         {object} handler.Response{data=GetProductBalanceResponse{}}
// @Failure     400         {object} handler.Response{error=handler.Error{}} "INVALID_ID"
// @Failure     500         {object} handler.Response{error=handler.Error{}} "ERR_INTERNAL_SERVER_ERROR"
// @Router      /menu/products/{product_id}/stock-balance [GET]
func (s Service) GetProductStockBalance(ctx context.Context, r *http.Request) handler.Response {
	id, err := handler.GetResourceIDFromURL(r, "product_id")
	if err != nil {
		return handler.IDParsingErrorResponse(err)
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
