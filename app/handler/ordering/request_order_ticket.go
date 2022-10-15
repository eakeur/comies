package ordering

import (
	"comies/app/handler/rest"
	"comies/app/workflows/ordering"
	"context"
	"net/http"
)

// RequestOrderTicket
// @Tags        Ordering
// @Success     201         {object} rest.Response{data=OrderRequestResponse{}}
// @Failure     500         {object} rest.Response{error=rest.Error{}} "ERR_INTERNAL_SERVER_ERROR"
// @Router      /ordering/new [POST]
func RequestOrderTicket(ctx context.Context, _ *http.Request) rest.Response {
	ticket, err := ordering.InitializeOrder(ctx)
	if err != nil {
		return rest.Fail(err)
	}

	return rest.ResponseWithData(http.StatusCreated, OrderRequestResponse{ID: ticket.String()})
}
