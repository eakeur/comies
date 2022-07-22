package ordering

import (
	"comies/app/gateway/api/response"
	"comies/app/sdk/throw"
	"context"
	"net/http"
)

func (s Service) RequestOrderTicket(ctx context.Context) response.Response {
	ticket, err := s.ordering.RequestOrderTicket(ctx)
	if err != nil {
		return failures.Handle(throw.Error(err))
	}

	return response.WithData(http.StatusCreated, AdditionResult{ID: ticket})
}
