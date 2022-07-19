package ordering

import (
	"comies/app/gateway/api/gen/ordering/protos"
	"comies/app/sdk/throw"
	"comies/app/sdk/types"
	"context"
)

func (s service) CancelOrder(ctx context.Context, request *protos.CancelOrderRequest) (*protos.Empty, error) {
	err := s.ordering.CancelOrder(ctx, types.ID(request.Id))
	if err != nil {
		return nil, failures.HandleError(throw.Error(err))
	}

	return &protos.Empty{}, nil
}
