package ordering

import (
	"comies/app/gateway/api/handler"
	"context"
	"net/http"
)

// RequestOrderTicket
// @Tags        Ordering
// @Success     201         {object} handler.Response{data=OrderRequestResponse{}}
// @Failure     500         {object} handler.Response{error=handler.Error{}} "ERR_INTERNAL_SERVER_ERROR"
// @Router      /ordering/new [POST]
func (s Service) RequestOrderTicket(ctx context.Context, _ *http.Request) handler.Response {
	ticket, err := s.ordering.RequestOrderTicket(ctx)
	if err != nil {
		return handler.Fail(err)
	}

	return handler.ResponseWithData(http.StatusCreated, OrderRequestResponse{ID: ticket.String()})
}
