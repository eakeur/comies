package ordering

import (
	"comies/app/core/workflows/ordering"
	"comies/app/gateway/api/response"
	"comies/app/sdk/throw"
	"context"
	"net/http"
)

func (s Service) Order(ctx context.Context, c ordering.OrderConfirmation) response.Response {
	o, err := s.ordering.Order(ctx, c)
	if err != nil {
		return failures.Handle(throw.Error(err))
	}

	return response.WithData(http.StatusCreated, Order{
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
