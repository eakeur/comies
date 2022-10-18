package ordering

import (
	"comies/app/api/request"
	"comies/app/api/send"
	"comies/app/workflows/ordering"
	"context"
	"encoding/json"
	"net/http"
)

// SetOrderDeliveryMode
// @Tags        Ordering
// @Param       order_id path     string                  false "The order ID"
// @Param       status  body     SetOrderDeliveryModeRequest true  "The properties defining the mode"
// @Success     204
// @Failure     400         {object} rest.Response{error=rest.Error{}} "INVALID_ID"
// @Failure     500         {object} rest.Response{error=rest.Error{}} "ERR_INTERNAL_SERVER_ERROR"
// @Router      /ordering/orders/{order_id}/delivery-mode [PUT]
func SetOrderDeliveryMode(ctx context.Context, r request.Request) send.Response {
	id, err := r.IDParam("order_id")
	if err != nil {
		return send.IDError(err)
	}

	var req SetOrderDeliveryTypeRequest
	err = json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return send.JSONError(err)
	}

	err = ordering.SetOrderDeliveryMode(ctx, id, req.Type)
	if err != nil {
		return send.FromError(err)
	}

	return send.Data(http.StatusNoContent, nil)
}
