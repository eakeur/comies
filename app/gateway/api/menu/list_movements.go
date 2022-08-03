package menu

import (
	"comies/app/core/entities/movement"
	"comies/app/gateway/api/failures"
	"comies/app/gateway/api/handler"
	"comies/app/sdk/throw"
	"context"
	"net/http"
	"time"
)

func (s Service) GetProductMovements(ctx context.Context, r *http.Request) handler.Response {
	id, err, res := handler.GetResourceIDFromURL(r, "product_id")
	if err != nil {
		return res
	}

	var filter movement.Filter
	filter.ProductID = id

	var query = r.URL.Query()
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

	return handler.ResponseWithData(http.StatusOK, movements)
}
