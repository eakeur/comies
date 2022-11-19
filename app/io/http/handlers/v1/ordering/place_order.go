package ordering

import (
	"comies/app/io/http/request"
	"comies/app/io/http/send"
	"comies/app/jobs/ordering"
	"context"
	"encoding/json"
	"net/http"
)

// PlaceOrder
// @Tags        Ordering
// @Param       order_id path     string                  false "The order ID"
// @Param       confirmation  body     ConfirmOrderRequest true  "The properties defining the confirmation"
// @Success     200         {object} rest.Response{data=PlaceOrder{}}
// @Failure     400         {object} rest.Response{error=rest.Error{}} "INVALID_ID"
// @Failure     500         {object} rest.Response{error=rest.Error{}} "ERR_INTERNAL_SERVER_ERROR"
// @Router      /ordering/orders/{order_id}/confirm [POST]
func PlaceOrder(ctx context.Context, r request.Request) send.Response {

	id, err := r.IDParam("order_id")
	if err != nil {
		return send.IDError(err)
	}

	var c ConfirmOrderRequest
	err = json.NewDecoder(r.Body).Decode(&c)
	if err != nil {
		return send.JSONError(err)
	}

	o, err := ordering.PlaceOrder(ctx, ordering.OrderConfirmation{
		OrderID:         id,
		DeliveryType:    c.DeliveryMode,
		CustomerName:    c.CustomerName,
		CustomerAddress: c.CustomerAddress,
		CustomerPhone:   c.CustomerPhone,
	})
	if err != nil {
		return send.FromError(err)
	}

	r.Commit(ctx)
	return send.Data(http.StatusCreated, o)
}

type ConfirmOrderRequest struct {
	DeliveryMode    ordering.Type `json:"delivery_type"`
	CustomerName    string        `json:"customer_name"`
	CustomerAddress string        `json:"customer_address"`
	CustomerPhone   string        `json:"customer_phone"`
}
