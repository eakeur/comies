package ordering

import (
	"comies/app/gateway/api/failures"
	"comies/app/gateway/api/handler"
	"comies/app/sdk/throw"
	"context"
	"net/http"
)

func (s Service) RequestOrderTicket(ctx context.Context) handler.Response {
	ticket, err := s.ordering.RequestOrderTicket(ctx)
	if err != nil {
		return failures.Handle(throw.Error(err))
	}

	return handler.ResponseWithData(http.StatusCreated, AdditionResult{ID: ticket})
}
