package orders

import (
	"comies/app/io/http/request"
	"comies/app/io/http/send"
	"context"
	"net/http"
)

func (h Handler) Cancel(ctx context.Context, r request.Request) send.Response {
	id, err := r.IDParam("order_id")
	if err != nil {
		return send.IDError(err)
	}

	err = h.orders.CancelOrder(ctx, id)
	if err != nil {
		return send.FromError(err)
	}

	r.Commit(ctx)
	return send.Data(http.StatusNoContent, nil)
}
