package ordering

import (
	"comies/app/gateway/api/gen/ordering/protos"
	"comies/app/sdk/throw"
	"comies/app/sdk/types"
	"context"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (s service) GetOrderById(ctx context.Context, request *protos.GetOrderByIdRequest) (*protos.GetOrderByIdResponse, error) {
	o, err := s.ordering.GetOrderByID(ctx, types.ID(request.Id))
	if err != nil {
		return nil, failures.HandleError(throw.Error(err))
	}

	return &protos.GetOrderByIdResponse{
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
