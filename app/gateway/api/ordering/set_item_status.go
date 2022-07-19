package ordering

import (
	"comies/app/core/entities/item"
	"comies/app/gateway/api/gen/ordering/protos"
	"comies/app/sdk/throw"
	"comies/app/sdk/types"
	"context"
)

func (s service) SetItemStatus(ctx context.Context, request *protos.SetItemStatusRequest) (*protos.Empty, error) {
	err := s.ordering.SetItemStatus(ctx, types.ID(request.ItemId), item.Status(request.Status))
	if err != nil {
		return nil, failures.HandleError(throw.Error(err))
	}

	return &protos.Empty{}, nil
}
