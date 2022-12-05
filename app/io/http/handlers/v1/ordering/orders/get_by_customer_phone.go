package orders

import (
	"comies/app/io/http/request"
	"comies/app/io/http/send"
	"context"
	"net/http"
)

func (h Handler) GetByCustomerPhone(ctx context.Context, r request.Request) send.Response {
	phone := r.Param("customer_phone")

	o, err := h.orders.GetOrderByCustomerPhone(ctx, phone)
	if err != nil {
		return send.FromError(err)
	}

	return send.Data(http.StatusOK, o)
}
