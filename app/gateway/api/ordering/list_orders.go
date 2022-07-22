package ordering

import (
	"comies/app/core/entities/order"
	"comies/app/gateway/api/response"
	"comies/app/sdk/throw"
	"context"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

func (s Service) ListOrders(ctx context.Context, query url.Values) response.Response {
	var filter order.Filter
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

	list, err := s.ordering.ListOrders(ctx, filter)
	if err != nil {
		return failures.Handle(throw.Error(err))
	}

	orders := make([]Order, len(list))
	for i, o := range list {
		orders[i] = Order{
			ID:             o.ID,
			Identification: o.Identification,
			PlacedAt:       o.PlacedAt,
			Status:         o.Status,
			DeliveryMode:   o.DeliveryMode,
			Observations:   o.Observations,
			FinalPrice:     o.FinalPrice,
			Address:        o.Address,
			Phone:          o.Phone,
		}
	}

	return response.WithData(http.StatusOK, list)
}
