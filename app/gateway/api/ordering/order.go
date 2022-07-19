package ordering

import (
	"comies/app/core/entities/order"
	"comies/app/core/workflows/ordering"
	"comies/app/gateway/api/gen/ordering/protos"
	"comies/app/sdk/throw"
	"comies/app/sdk/types"
	"context"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (s service) Order(ctx context.Context, request *protos.OrderRequest) (*protos.OrderResponse, error) {
	o, err := s.ordering.Order(ctx, ordering.OrderConfirmation{
		OrderID:      types.ID(request.OrderId),
		DeliveryMode: order.DeliveryMode(request.DeliveryMode),
	})
	if err != nil {
		return nil, failures.HandleError(throw.Error(err))
	}

	return &protos.OrderResponse{
		Order: &protos.Order{
			Id:             int64(o.ID),
			Identification: o.Identification,
			PlacedAt:       timestamppb.New(o.PlacedAt),
			Status:         protos.OrderStatus(o.Status),
			DeliveryMode:   protos.DeliveryMode(o.DeliveryMode),
			Observation:    o.Observations,
			FinalPrice:     int64(o.FinalPrice),
			Address:        o.Address,
			Phone:          o.Phone,
		},
	}, nil
}
