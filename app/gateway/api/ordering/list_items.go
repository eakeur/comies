package ordering

import (
	"comies/app/gateway/api/gen/ordering/protos"
	"comies/app/sdk/throw"
	"comies/app/sdk/types"
	"context"
)

func (s service) ListItems(ctx context.Context, request *protos.ListItemsRequest) (*protos.ListItemsResponse, error) {
	items, err := s.ordering.ListItems(ctx, types.ID(request.OrderId))
	if err != nil {
		return nil, failures.HandleError(throw.Error(err))
	}

	list := make([]*protos.Item, len(items))
	for i, it := range items {
		ignore := make([]int64, len(it.Details.IgnoreIngredients))
		replacements := make(map[int64]int64, len(it.Details.ReplaceIngredients))
		for i, id := range it.Details.IgnoreIngredients {
			ignore[i] = int64(id)
		}
		for from, to := range it.Details.ReplaceIngredients {
			replacements[int64(from)] = int64(to)
		}

		list[i] = &protos.Item{
			Id:           int64(it.ID),
			OrderId:      int64(it.OrderID),
			ProductId:    int64(it.ProductID),
			Price:        int64(it.Price),
			Status:       protos.ItemStatus(it.Status),
			Quantity:     int64(it.Quantity),
			Observations: it.Observations,
			Ignored:      ignore,
			Replacements: replacements,
		}
	}

	return &protos.ListItemsResponse{
		Items: list,
	}, nil

}
