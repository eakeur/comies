package ordering

import (
	"comies/app/core/id"
	"comies/app/core/order"
	"comies/app/core/types"
	"comies/app/handler/rest"
	"comies/app/workflows/ordering"
	"context"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type ListOrdersResponse struct {
	ID             id.ID              `json:"id"`
	Identification string             `json:"identification,omitempty"`
	PlacedAt       time.Time          `json:"placed_at"`
	Status         order.Status       `json:"status,omitempty"`
	DeliveryMode   order.DeliveryMode `json:"delivery_mode,omitempty"`
	Observations   string             `json:"observations,omitempty"`
	FinalPrice     types.Currency     `json:"final_price,omitempty"`
	Address        string             `json:"address,omitempty"`
	Phone          string             `json:"phone,omitempty"`
	Items          []Item             `json:"items,omitempty"`
}

// ListOrders
//
// @Tags        Ordering
// @Success     200         {object} rest.Response{data=[]PlaceOrder{}}
// @Failure     500         {object} rest.Response{error=rest.Error{}} "ERR_INTERNAL_SERVER_ERROR"
// @Router      /ordering/ordering [GET]
func ListOrders(ctx context.Context, r *http.Request) rest.Response {
	var filter order.Filter

	query := r.URL.Query()
	if sts := strings.Split(query.Get("status"), ","); len(sts) > 0 {
		filter.Status = make([]order.Status, len(sts))
		for i, o := range sts {
			s, err := strconv.Atoi(o)
			if err != nil {
				continue
			}
			filter.Status[i] = order.Status(s)
		}
	}
	if parse, err := time.Parse(time.RFC3339, query.Get("before")); err == nil {
		filter.PlacedBefore = parse
	}

	if parse, err := time.Parse(time.RFC3339, query.Get("after")); err == nil {
		filter.PlacedAfter = parse
	}

	list, err := ordering.ListOrders(ctx, filter)
	if err != nil {
		return rest.Fail(err)
	}

	return rest.ResponseWithData(http.StatusOK, NewOrderList(list))
}
