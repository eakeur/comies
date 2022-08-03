package ordering

import (
	"comies/app/core/workflows/ordering"
	"comies/app/gateway/api/failures"
	"comies/app/gateway/api/handler"
	"comies/app/sdk/throw"
	"context"
	"encoding/json"
	"net/http"
)

func (s Service) Order(ctx context.Context, r *http.Request) handler.Response {

	var c ordering.OrderConfirmation
	err := json.NewDecoder(r.Body).Decode(&c)
	if err != nil {
		return handler.JSONParsingErrorResponse(err)
	}

	o, err := s.ordering.Order(ctx, c)
	if err != nil {
		return failures.Handle(throw.Error(err))
	}

	return handler.ResponseWithData(http.StatusCreated, Order{
		ID:             o.ID,
		Identification: o.Identification,
		PlacedAt:       o.PlacedAt,
		Status:         o.Status,
		DeliveryMode:   o.DeliveryMode,
		Observations:   o.Observations,
		FinalPrice:     o.FinalPrice,
		Address:        o.Address,
		Phone:          o.Phone,
	})
}
