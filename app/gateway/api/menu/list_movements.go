package menu

import (
	"comies/app/core/entities/movement"
	"comies/app/gateway/api/handler"
	"comies/app/gateway/api/response"
	"comies/app/sdk/throw"
	"context"
	"net/http"
	"net/url"
	"time"
)

func (s Service) GetProductMovements(ctx context.Context, params handler.RouteParams, query url.Values) response.Response {
	id, err, res := convertToID(params["product_id"])
	if err != nil {
		return res
	}

	var filter movement.Filter
	filter.ProductID = id

	if parse, err := time.Parse(time.RFC3339, query.Get("start")); err == nil {
		filter.InitialDate = parse
	}

	if parse, err := time.Parse(time.RFC3339, query.Get("end")); err == nil {
		filter.FinalDate = parse
	}

	list, err := s.menu.ListMovements(ctx, filter)
	if err != nil {
		return failures.Handle(throw.Error(err))
	}

	movements := make([]Movement, len(list))
	for i, p := range list {
		movements[i] = Movement{
			ID:        p.ID.String(),
			ProductID: p.ProductID.String(),
			Type:      p.Type,
			Date:      p.Date,
			Quantity:  p.Quantity,
			PaidValue: p.PaidValue,
			AgentID:   p.AgentID.String(),
		}
	}

	return response.WithData(http.StatusOK, movements)
}
