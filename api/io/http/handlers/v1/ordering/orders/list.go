package orders

import (
	"comies/core/ordering/order"
	"comies/core/types"
	"comies/io/http/request"
	"comies/io/http/send"
	"context"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func (h Handler) List(ctx context.Context, r request.Request) send.Response {
	var filter order.Filter

	query := r.URL.Query()
	if sts := strings.Split(query.Get("status"), ","); len(sts) > 0 {
		filter.Status = make([]types.Status, len(sts))
		for i, o := range sts {
			s, err := strconv.Atoi(o)
			if err != nil {
				continue
			}
			filter.Status[i] = types.Status(s)
		}
	}
	if parse, err := time.Parse(time.RFC3339, query.Get("before")); err == nil {
		filter.PlacedBefore = parse
	}

	if parse, err := time.Parse(time.RFC3339, query.Get("after")); err == nil {
		filter.PlacedAfter = parse
	}

	o, err := h.orders.ListOrders(ctx, filter)
	if err != nil {
		return send.FromError(err)
	}

	return send.Data(http.StatusOK, o)
}
