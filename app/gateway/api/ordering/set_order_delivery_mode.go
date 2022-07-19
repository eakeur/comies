package ordering

import (
	"comies/app/core/entities/order"
	"comies/app/gateway/api/gen/ordering/protos"
	"comies/app/sdk/throw"
	"comies/app/sdk/types"
	"context"
)

func (s service) SetOrderDeliveryMode(ctx context.Context, request *protos.SetOrderDeliveryModeRequest) (*protos.Empty, error) {
	err := s.ordering.SetOrderDeliveryMode(ctx, types.ID(request.OrderId), order.DeliveryMode(request.DeliveryMode))
	if err != nil {
		return nil, failures.HandleError(throw.Error(err))
	}

	return &protos.Empty{}, nil
}
