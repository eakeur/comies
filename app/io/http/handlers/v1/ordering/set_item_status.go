package ordering

import (
	"comies/app/io/http/request"
	"comies/app/io/http/send"
	"comies/app/jobs/ordering"
	"context"
	"encoding/json"
	"net/http"
)

// SetItemStatus
// @Tags        Ordering
// @Param       order_id path     string                  false "The order ID"
// @Param       status  body     SetItemStatusRequest true  "The properties defining the status"
// @Success     204
// @Failure     400         {object} rest.Response{error=rest.Error{}} "INVALID_ID"
// @Failure     500         {object} rest.Response{error=rest.Error{}} "ERR_INTERNAL_SERVER_ERROR"
// @Router      /ordering/orders/{order_id}/items/{item_id} [PUT]
func SetItemStatus(ctx context.Context, r request.Request) send.Response {
	id, err := r.IDParam("item_id")
	if err != nil {
		return send.IDError(err)
	}

	var req SetItemStatusRequest
	err = json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return send.JSONError(err)
	}

	err = ordering.SetItemStatus(ctx, id, req.Status)
	if err != nil {
		return send.FromError(err)
	}

	r.Commit(ctx)
	return send.Data(http.StatusNoContent, nil)
}

type SetItemStatusRequest struct {
	Status ordering.Status `json:"status"`
}
