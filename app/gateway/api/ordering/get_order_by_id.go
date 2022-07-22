package ordering

import (
	"comies/app/gateway/api/handler"
	"comies/app/gateway/api/response"
	"comies/app/sdk/throw"
	"context"
	"net/http"
)

func (s Service) GetOrderById(ctx context.Context, params handler.RouteParams) response.Response {
	id, err, res := convertToID(params["order_id"])
	if err != nil {
		return res
	}

	o, err := s.ordering.GetOrderByID(ctx, id)
	if err != nil {
		return failures.Handle(throw.Error(err))
	}

	return response.WithData(http.StatusOK, Order{
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
