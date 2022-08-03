package ordering

import (
	"comies/app/gateway/api/failures"
	"comies/app/gateway/api/handler"
	"comies/app/sdk/throw"
	"context"
	"net/http"
)

func (s Service) GetOrderByID(ctx context.Context, r *http.Request) handler.Response {
	id, err, res := handler.GetResourceIDFromURL(r, "order_id")
	if err != nil {
		return res
	}

	o, err := s.ordering.GetOrderByID(ctx, id)
	if err != nil {
		return failures.Handle(throw.Error(err))
	}

	return handler.ResponseWithData(http.StatusOK, Order{
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
