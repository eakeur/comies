package ordering

import (
	"comies/app/core/workflows/ordering"
	"comies/app/gateway/api/handler"
	"context"
	"encoding/json"
	"net/http"
)

// Order
// @Tags        Ordering
// @Param       order_id path     string                  false "The order ID"
// @Param       confirmation  body     ConfirmOrderRequest true  "The properties defining the confirmation"
// @Success     200         {object} handler.Response{data=Order{}}
// @Failure     400         {object} handler.Response{error=handler.Error{}} "INVALID_ID"
// @Failure     500         {object} handler.Response{error=handler.Error{}} "ERR_INTERNAL_SERVER_ERROR"
// @Router      /ordering/orders/{order_id}/confirm [POST]
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
		return handler.Fail(err)
	}

	return handler.ResponseWithData(http.StatusCreated, NewOrder(o))
}
