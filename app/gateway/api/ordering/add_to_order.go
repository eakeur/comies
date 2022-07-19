package ordering

import (
	"comies/app/core/entities/item"
	"comies/app/gateway/api/gen/ordering/protos"
	"comies/app/sdk/throw"
	"comies/app/sdk/types"
	"context"
)

func (s service) AddToOrder(ctx context.Context, request *protos.AddToOrderRequest) (*protos.AddToOrderResponse, error) {
	ignore := make([]types.ID, len(request.Item.Ignored))
	replacements := make(map[types.ID]types.ID, len(request.Item.Replacements))
	for i, id := range request.Item.Ignored {
		ignore[i] = types.ID(id)
	}
	for from, to := range request.Item.Replacements {
		replacements[types.ID(from)] = types.ID(to)
	}

	it, err := s.ordering.AddToOrder(ctx, item.Item{
		OrderID:      types.ID(request.Item.OrderId),
		ProductID:    types.ID(request.Item.ProductId),
		Quantity:     types.Quantity(request.Item.Quantity),
		Observations: request.Item.Observations,
		Details: item.Details{
			ReplaceIngredients: replacements,
			IgnoreIngredients:  ignore,
		},
	})
	if err != nil {
		return nil, failures.HandleError(throw.Error(err))
	}

	failures := make([]*protos.ReservationFailure, len(it.Failed))
	for _, f := range it.Failed {
		failures = append(failures, &protos.ReservationFailure{
			ProductId: int64(f.ProductID),
			Error:     f.Error.Error(),
		})
	}

	return &protos.AddToOrderResponse{Failures: failures}, nil
}
