package orders

import (
	"comies/app/core/ordering/status"
	"comies/app/core/types"
	"comies/app/io/http/request"
	"comies/app/io/http/send"
	"context"
	"net/http"
	"time"
)

func (h Handler) SetStatus(ctx context.Context, r request.Request) send.Response {
	id, err := r.IDParam("order_id")
	if err != nil {
		return send.IDError(err)
	}

	var st SetOrderStatusRequest
	err = r.JSONBody(&st)
	if err != nil {
		return send.JSONError(err)
	}

	err = h.orders.SetOrderStatus(ctx, id, status.Status{
		OrderID:    id,
		OccurredAt: time.Now(),
		Value:      st.Status,
	})
	if err != nil {
		return send.FromError(err)
	}

	r.Commit(ctx)
	return send.Data(http.StatusNoContent, nil)
}

type SetOrderStatusRequest struct {
	Status types.Status `json:"status"`
}
