package ordering

import (
	"comies/app/core/workflows/ordering"
	"comies/app/handler/rest"
	"context"
	"encoding/json"
	"net/http"
)

// Order
// @Tags        Ordering
// @Param       order_id path     string                  false "The order ID"
// @Param       confirmation  body     ConfirmOrderRequest true  "The properties defining the confirmation"
// @Success     200         {object} rest.Response{data=Order{}}
// @Failure     400         {object} rest.Response{error=rest.Error{}} "INVALID_ID"
// @Failure     500         {object} rest.Response{error=rest.Error{}} "ERR_INTERNAL_SERVER_ERROR"
// @Router      /ordering/orders/{order_id}/confirm [POST]
func PlaceOrder(ctx context.Context, r *http.Request) rest.Response {

	id, err := rest.GetResourceIDFromURL(r, "order_id")
	if err != nil {
		return rest.IDParsingErrorResponse(err)
	}

	var c ConfirmOrderRequest
	err = json.NewDecoder(r.Body).Decode(&c)
	if err != nil {
		return rest.JSONParsingErrorResponse(err)
	}

	o, err := ordering.Order(ctx, ordering.OrderConfirmation{
		OrderID:      id,
		DeliveryMode: c.DeliveryMode,
	})
	if err != nil {
		return rest.Fail(err)
	}

	return rest.ResponseWithData(http.StatusCreated, NewOrder(o))
}
