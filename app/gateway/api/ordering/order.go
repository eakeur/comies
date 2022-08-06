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

	id, err := handler.GetResourceIDFromURL(r, "order_id")
	if err != nil {
		return handler.IDParsingErrorResponse(err)
	}

	var c ConfirmOrderRequest
	err = json.NewDecoder(r.Body).Decode(&c)
	if err != nil {
		return handler.JSONParsingErrorResponse(err)
	}

	o, err := s.ordering.Order(ctx, ordering.OrderConfirmation{
		OrderID:      id,
		DeliveryMode: c.DeliveryMode,
	})
	if err != nil {
		return failures.Handle(throw.Error(err))
	}

	return handler.ResponseWithData(http.StatusCreated, NewOrder(o))
}
