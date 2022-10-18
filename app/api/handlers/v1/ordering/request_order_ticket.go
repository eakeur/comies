package ordering

import (
	"comies/app/api/request"
	"comies/app/api/send"
	"comies/app/jobs/ordering"
	"context"
)

// RequestOrderTicket
// @Tags        Ordering
// @Success     201         {object} rest.Response{data=OrderRequestResponse{}}
// @Failure     500         {object} rest.Response{error=rest.Error{}} "ERR_INTERNAL_SERVER_ERROR"
// @Router      /ordering/new [POST]
func RequestOrderTicket(ctx context.Context, _ request.Request) send.Response {
	ticket, err := ordering.InitializeOrder(ctx)
	if err != nil {
		return send.FromError(err)
	}

	return send.CreatedWithID(ticket.ID)
}
