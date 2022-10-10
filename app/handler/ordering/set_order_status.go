package ordering

import (
	"comies/app/core/workflows/ordering"
	"comies/app/handler/rest"
	"context"
	"encoding/json"
	"net/http"
)

// SetOrderStatus
// @Tags        Ordering
// @Param       order_id path     string                  false "The order ID"
// @Param       status  body     SetOrderStatus true  "The properties defining the order status"
// @Success     204
// @Failure     400         {object} rest.Response{error=rest.Error{}} "INVALID_ID"
// @Failure     500         {object} rest.Response{error=rest.Error{}} "ERR_INTERNAL_SERVER_ERROR"
// @Router      /ordering/orders/{order_id}/status [PUT]
func SetOrderStatus(ctx context.Context, r *http.Request) rest.Response {
	id, err := rest.GetResourceIDFromURL(r, "order_id")
	if err != nil {
		return rest.IDParsingErrorResponse(err)
	}

	var req SetOrderStatusRequest
	err = json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return rest.JSONParsingErrorResponse(err)
	}

	err = ordering.SetOrderStatus(ctx, id, req.Status)
	if err != nil {
		return rest.Fail(err)
	}

	return rest.ResponseWithData(http.StatusNoContent, nil)
}
