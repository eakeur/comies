package menu

import (
	"comies/app/core/entities/movement"
	"comies/app/gateway/api/handler"
	"comies/app/gateway/api/response"
	"comies/app/sdk/throw"
	"context"
	"net/http"
)

func (s Service) GetProductBalance(ctx context.Context, params handler.RouteParams) response.Response {
	id, err, res := convertToID(params["product_id"])
	if err != nil {
		return res
	}

	bal, err := s.menu.GetMovementsBalance(ctx, movement.Filter{
		ProductID: id,
	})
	if err != nil {
		return failures.Handle(throw.Error(err))
	}

	return response.WithData(http.StatusOK, ProductStockBalanceResult{Balance: bal})
}
