package ordering

import (
	"comies/app/gateway/api/handler"
	"context"
	"encoding/json"
	"net/http"
)

// SetOrderDeliveryMode
// @Tags        Ordering
// @Param       order_id path     string                  false "The order ID"
// @Param       status  body     SetOrderDeliveryModeRequest true  "The properties defining the mode"
// @Success     204
// @Failure     400         {object} handler.Response{error=handler.Error{}} "INVALID_ID"
// @Failure     500         {object} handler.Response{error=handler.Error{}} "ERR_INTERNAL_SERVER_ERROR"
// @Router      /ordering/orders/{order_id}/delivery-mode [PUT]
func (s Service) SetOrderDeliveryMode(ctx context.Context, r *http.Request) handler.Response {
	id, err := handler.GetResourceIDFromURL(r, "order_id")
	if err != nil {
		return handler.IDParsingErrorResponse(err)
	}

	var req SetOrderDeliveryModeRequest
	err = json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return handler.JSONParsingErrorResponse(err)
	}

	err = s.ordering.SetOrderDeliveryMode(ctx, id, req.Mode)
	if err != nil {
		return handler.Fail(err)
	}

	return handler.ResponseWithData(http.StatusNoContent, nil)
}
