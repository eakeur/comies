package statuses

import (
	"comies/api/request"
	"comies/api/send"
	"comies/core/ordering/status"
	"comies/core/types"
	"context"
	"net/http"
	"strconv"
	"time"
)

func (h Handler) Set(ctx context.Context, r request.Request) send.Response {
	id, err := r.IDParam("order_id")
	if err != nil {
		return send.IDError(err)
	}

	st, err := strconv.Atoi(r.Param("status"))
	if err != nil {
		return send.Status(http.StatusBadRequest)
	}

	err = h.statuses.SetOrderStatus(ctx, id, status.Status{
		OrderID:    id,
		OccurredAt: time.Now(),
		Value:      types.Status(st),
	})
	if err != nil {
		return send.FromError(err)
	}

	r.Commit(ctx)
	return send.Data(http.StatusNoContent, nil)
}
