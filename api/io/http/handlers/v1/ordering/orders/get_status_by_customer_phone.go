package orders

import (
	"comies/io/http/request"
	"comies/io/http/send"
	"context"
	"net/http"
)

func (h Handler) GetByCustomerPhone(ctx context.Context, r request.Request) send.Response {
	phone := r.Param("customer_phone")

	o, err := h.orders.GetStatusByCustomerPhone(ctx, phone)
	if err != nil {
		return send.FromError(err)
	}

	return send.Data(http.StatusOK, o)
}
