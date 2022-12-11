package orders

import (
	"comies/core/types"
	"comies/io/http/request"
	"comies/io/http/send"
	"context"
	"net/http"
)

func (h Handler) GetCustomer(ctx context.Context, r request.Request) send.Response {
	key := r.Param("order_id")

	if r.GetQuery("phone") == "true" {
		o, err := h.orders.GetStatusByCustomerPhone(ctx, key)
		if err != nil {
			return send.FromError(err)
		}

		return send.Data(http.StatusOK, o)
	}

	id, err := types.FromString(key)
	if err != nil {
		return send.IDError(err)
	}

	o, err := h.orders.GetOrderByID(ctx, id)
	if err != nil {
		return send.FromError(err)
	}

	return send.Data(http.StatusOK, o)
}
