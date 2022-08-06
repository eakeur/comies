package ordering

import (
	"comies/app/core/entities/order"
	"comies/app/gateway/api/failures"
	"comies/app/gateway/api/handler"
	"comies/app/sdk/throw"
	"comies/app/sdk/types"
	"context"
	"net/http"
	"time"
)

type GetOrderByIDResponse struct {
	ID             types.ID           `json:"id"`
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

func (s Service) GetOrderByID(ctx context.Context, r *http.Request) handler.Response {
	id, err := handler.GetResourceIDFromURL(r, "order_id")
	if err != nil {
		return handler.IDParsingErrorResponse(err)
	}

	o, err := s.ordering.GetOrderByID(ctx, id)
	if err != nil {
		return failures.Handle(throw.Error(err))
	}

	return handler.ResponseWithData(http.StatusOK, NewGetOrderByIDResponse(o))
}

func NewGetOrderByIDResponse(o order.Order) GetOrderByIDResponse {
	return GetOrderByIDResponse{
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
