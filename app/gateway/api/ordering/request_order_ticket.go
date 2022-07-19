package ordering

import (
	"comies/app/gateway/api/gen/ordering/protos"
	"comies/app/sdk/throw"
	"context"
)

func (s service) RequestOrderTicket(ctx context.Context, _ *protos.Empty) (*protos.RequestOrderTicketResponse, error) {
	ticket, err := s.ordering.RequestOrderTicket(ctx)
	if err != nil {
		return nil, failures.HandleError(throw.Error(err))
	}

	return &protos.RequestOrderTicketResponse{OrderId: int64(ticket)}, nil
}
