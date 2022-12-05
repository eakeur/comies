package orders

import (
	"comies/app/io/http/request"
	"comies/app/io/http/send"
	"context"
	"net/http"
)

func (h Handler) GetByID(ctx context.Context, r request.Request) send.Response {
	id, err := r.IDParam("order_id")
	if err != nil {
		return send.IDError(err)
	}

	o, err := h.orders.GetOrderByID(ctx, id)
	if err != nil {
		return send.FromError(err)
	}

	return send.Data(http.StatusOK, o)
}
