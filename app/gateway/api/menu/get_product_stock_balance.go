package menu

import (
	"comies/app/core/entities/movement"
	"comies/app/gateway/api/gen/menu"
	"comies/app/sdk/throw"
	"comies/app/sdk/types"
	"context"
)

func (s service) GetProductBalance(ctx context.Context, in *menu.GetProductBalanceRequest) (*menu.GetProductBalanceResponse, error) {
	bal, err := s.menu.GetMovementsBalance(ctx, movement.Filter{
		ProductID: types.ID(in.Id),
	})
	if err != nil {
		return nil, failures.HandleError(throw.Error(err))
	}

	return &menu.GetProductBalanceResponse{
		Balance: int64(bal),
	}, nil
}
