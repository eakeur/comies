package menu

import (
	"comies/app/core/entities/movement"
	"comies/app/gateway/api/failures"
	"comies/app/gateway/api/handler"
	"comies/app/sdk/throw"
	"context"
	"net/http"
)

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

	return handler.ResponseWithData(http.StatusOK, ProductStockBalanceResult{Balance: bal})
}
