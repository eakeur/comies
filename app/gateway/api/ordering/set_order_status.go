package ordering

import (
	"comies/app/core/throw"
	"comies/app/gateway/api/handler"
	"context"
	"encoding/json"
	"net/http"
)

// SetOrderStatus
// @Tags        Ordering
// @Param       order_id path     string                  false "The order ID"
// @Param       status  body     SetOrderStatus true  "The properties defining the order status"
// @Success     204
// @Failure     400         {object} handler.Response{error=handler.Error{}} "INVALID_ID"
// @Failure     500         {object} handler.Response{error=handler.Error{}} "ERR_INTERNAL_SERVER_ERROR"
// @Router      /ordering/orders/{order_id}/status [PUT]
func (s Service) SetOrderStatus(ctx context.Context, r *http.Request) handler.Response {
	id, err := handler.GetResourceIDFromURL(r, "order_id")
	if err != nil {
		return handler.IDParsingErrorResponse(err)
	}

	var req SetOrderStatusRequest
	err = json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return handler.JSONParsingErrorResponse(err)
	}

	err = s.ordering.SetOrderStatus(ctx, id, req.Status)
	if err != nil {
		return handler.Fail(throw.Error(err))
	}

	return handler.ResponseWithData(http.StatusNoContent, nil)
}
