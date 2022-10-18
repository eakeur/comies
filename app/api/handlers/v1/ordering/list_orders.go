package ordering

import (
	"comies/app/api/request"
	"comies/app/api/send"
	"comies/app/data/orders"
	"comies/app/workflows/ordering"
	"context"
	"net/http"
	"strconv"
	"strings"
	"time"
)

// ListOrders
//
// @Tags        Ordering
// @Success     200         {object} rest.Response{data=[]PlaceOrder{}}
// @Failure     500         {object} rest.Response{error=rest.Error{}} "ERR_INTERNAL_SERVER_ERROR"
// @Router      /ordering/ordering [GET]
func ListOrders(ctx context.Context, r request.Request) send.Response {
	var filter ordering.OrderFilter

	query := r.URL.Query()
	if sts := strings.Split(query.Get("status"), ","); len(sts) > 0 {
		filter.Status = make([]ordering.Status, len(sts))
		for i, o := range sts {
			s, err := strconv.Atoi(o)
			if err != nil {
				continue
			}
			filter.Status[i] = ordering.Status(s)
		}
	}
	if parse, err := time.Parse(time.RFC3339, query.Get("before")); err == nil {
		filter.PlacedBefore = parse
	}

	if parse, err := time.Parse(time.RFC3339, query.Get("after")); err == nil {
		filter.PlacedAfter = parse
	}

	list, err := orders.List(ctx, filter)
	if err != nil {
		return send.FromError(err)
	}

	return send.Data(http.StatusOK, list)
}
