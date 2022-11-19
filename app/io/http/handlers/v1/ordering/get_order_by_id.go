package ordering

import (
	"comies/app/data/orders"
	"comies/app/io/http/request"
	"comies/app/io/http/send"
	"context"
	"net/http"
)

// GetOrderByID fetches a specific order
//
// @Summary     Get order by ID
// @Description Fetches an order looking for its ID
// @Tags        Ordering
// @Param       order_id path     string                  false "The order ID"
// @Success     200         {object} rest.Response{data=PlaceOrder{}}
// @Failure     400         {object} rest.Response{error=rest.Error{}} "INVALID_ID"
// @Failure     404         {object} rest.Response{data=[]Failure{}} "ORDER_NOT_FOUND"
// @Failure     500         {object} rest.Response{error=rest.Error{}} "ERR_INTERNAL_SERVER_ERROR"
// @Router      /ordering/orders/{order_id}/items [POST]
func GetOrderByID(ctx context.Context, r request.Request) send.Response {
	id, err := r.IDParam("order_id")
	if err != nil {
		return send.IDError(err)
	}

	o, err := orders.GetByID(ctx, id)
	if err != nil {
		return send.FromError(err)
	}

	return send.Data(http.StatusOK, o)
}
