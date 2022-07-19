package ordering

import (
	"comies/app/gateway/api/gen/ordering/protos"
	"comies/app/sdk/throw"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (s service) ListOrdersInFlow(_ *protos.Empty, server protos.Ordering_ListOrdersInFlowServer) error {
	channel, err := s.ordering.Channel(server.Context())
	if err != nil {
		return failures.HandleError(throw.Error(err))
	}

	for {
		select {
		case not := <-channel:
			items := make([]*protos.Item, len(not.Items))
			for ind, it := range not.Items {
				ignore := make([]int64, len(it.Details.IgnoreIngredients))
				replacements := make(map[int64]int64, len(it.Details.ReplaceIngredients))
				for i, id := range it.Details.IgnoreIngredients {
					ignore[i] = int64(id)
				}
				for from, to := range it.Details.ReplaceIngredients {
					replacements[int64(from)] = int64(to)
				}

				items[ind] = &protos.Item{
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
			err := server.Send(&protos.ListOrdersInFlowResponse{
				Order: &protos.Order{
					Id:             int64(not.ID),
					Identification: not.Identification,
					PlacedAt:       timestamppb.New(not.PlacedAt),
					Observation:    not.Observations,
					FinalPrice:     int64(not.FinalPrice),
					Address:        not.Address,
					Phone:          not.Phone,
				},
				Items: items,
			})
			if err != nil {
				return failures.HandleError(throw.Error(err))
			}
		}
	}
}
