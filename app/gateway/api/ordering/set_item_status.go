package ordering

import (
	"comies/app/gateway/api/handler"
	"context"
	"encoding/json"
	"net/http"
)

// SetItemStatus
// @Tags        Ordering
// @Param       order_id path     string                  false "The order ID"
// @Param       status  body     SetItemStatusRequest true  "The properties defining the status"
// @Success     204
// @Failure     400         {object} handler.Response{error=handler.Error{}} "INVALID_ID"
// @Failure     500         {object} handler.Response{error=handler.Error{}} "ERR_INTERNAL_SERVER_ERROR"
// @Router      /ordering/orders/{order_id}/items/{item_id} [PUT]
func (s Service) SetItemStatus(ctx context.Context, r *http.Request) handler.Response {
	id, err := handler.GetResourceIDFromURL(r, "item_id")
	if err != nil {
		return handler.IDParsingErrorResponse(err)
	}

	var req SetItemStatusRequest
	err = json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return handler.JSONParsingErrorResponse(err)
	}

	err = s.ordering.SetItemStatus(ctx, id, req.Status)
	if err != nil {
		return handler.Fail(err)
	}

	return handler.ResponseWithData(http.StatusNoContent, nil)
}
