package ordering

import (
	"comies/app/core/entities/order"
	"comies/app/gateway/api/gen/ordering/protos"
	"comies/app/sdk/throw"
	"comies/app/sdk/types"
	"context"
)

func (s service) SetOrderStatus(ctx context.Context, request *protos.SetOrderStatusRequest) (*protos.Empty, error) {
	err := s.ordering.SetOrderStatus(ctx, types.ID(request.OrderId), order.Status(request.Status))
	if err != nil {
		return nil, failures.HandleError(throw.Error(err))
	}

	return &protos.Empty{}, nil
}
