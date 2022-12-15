package orders

import (
	"comies/io/http/request"
	"comies/io/http/send"
	"comies/jobs/ordering"
	"context"
	"time"
)

func (h Handler) Place(ctx context.Context, r request.Request) send.Response {
	var order ordering.Order
	err := r.JSONBody(&order)
	if err != nil {
		return send.JSONError(err)
	}

	order.Date = time.Now()

	o, err := h.orders.PlaceOrder(ctx, order)
	if err != nil {
		return send.FromError(err)
	}

	r.Commit(ctx)

	return send.CreatedWithID(o.ID)
}
