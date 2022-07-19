package ordering

import (
	"comies/app/core/entities/order"
	"comies/app/gateway/api/gen/ordering/protos"
	"comies/app/sdk/throw"
	"context"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (s service) ListOrders(ctx context.Context, request *protos.ListOrdersRequest) (*protos.ListOrdersResponse, error) {
	var statuses []order.Status
	if l := len(request.Filter.Statuses); l > 0 {
		statuses = make([]order.Status, l)
		for i, o := range request.Filter.Statuses {
			statuses[i] = order.Status(o)
		}
	}
	list, err := s.ordering.ListOrders(ctx, order.Filter{
		Status:       statuses,
		PlacedBefore: request.Filter.PlacedBefore.AsTime(),
		PlacedAfter:  request.Filter.PlacedAfter.AsTime(),
		DeliveryMode: order.DeliveryMode(request.Filter.DeliveryMode),
	})
	if err != nil {
		return nil, failures.HandleError(throw.Error(err))
	}

	orders := make([]*protos.Order, len(list))
	for i, o := range list {
		orders[i] = &protos.Order{
			Id:             int64(o.ID),
			Identification: o.Identification,
			PlacedAt:       timestamppb.New(o.PlacedAt),
			Status:         protos.OrderStatus(o.Status),
			DeliveryMode:   protos.DeliveryMode(o.DeliveryMode),
			Observation:    o.Observations,
			FinalPrice:     int64(o.FinalPrice),
			Address:        o.Address,
			Phone:          o.Phone,
		}
	}

	return &protos.ListOrdersResponse{
		Orders: orders,
	}, nil
}
