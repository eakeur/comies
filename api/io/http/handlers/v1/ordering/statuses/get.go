package statuses

import (
	"comies/io/http/request"
	"comies/io/http/send"
	"context"
	"net/http"
)

func (h Handler) Get(ctx context.Context, r request.Request) send.Response {
	key := r.Param("order_id")

	if r.GetQuery("phone") == "true" {
		o, err := h.statuses.GetStatusByCustomerPhone(ctx, key)
		if err != nil {
			return send.FromError(err)
		}

		return send.Data(http.StatusOK, o)
	}

	return send.Status(http.StatusNotImplemented)
}
