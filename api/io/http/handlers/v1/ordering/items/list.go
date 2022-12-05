package items

import (
	"comies/io/http/request"
	"comies/io/http/send"
	"context"
	"net/http"
)

func (h Handler) List(ctx context.Context, r request.Request) send.Response {
	orderID, err := r.IDParam("order_id")
	if err != nil {
		return send.IDError(err)
	}

	o, err := h.items.ListItems(ctx, orderID)
	if err != nil {
		return send.FromError(err)
	}

	return send.Data(http.StatusOK, o)
}
