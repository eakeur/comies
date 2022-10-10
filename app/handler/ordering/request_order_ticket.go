package ordering

import (
	"comies/app/core/workflows/ordering"
	"comies/app/handler/rest"
	"context"
	"net/http"
)

// RequestOrderTicket
// @Tags        Ordering
// @Success     201         {object} rest.Response{data=OrderRequestResponse{}}
// @Failure     500         {object} rest.Response{error=rest.Error{}} "ERR_INTERNAL_SERVER_ERROR"
// @Router      /ordering/new [POST]
func RequestOrderTicket(ctx context.Context, _ *http.Request) rest.Response {
	ticket, err := ordering.RequestOrderTicket(ctx)
	if err != nil {
		return rest.Fail(err)
	}

	return rest.ResponseWithData(http.StatusCreated, OrderRequestResponse{ID: ticket.String()})
}
