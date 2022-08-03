package ordering

import (
	"comies/app/gateway/api/failures"
	"comies/app/gateway/api/handler"
	"comies/app/sdk/throw"
	"context"
	"encoding/json"
	"net/http"
)

func (s Service) SetItemStatus(ctx context.Context, r *http.Request) handler.Response {
	id, err, res := handler.GetResourceIDFromURL(r, "order_id")
	if err != nil {
		return res
	}

	var req SetItemStatusRequest
	err = json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return handler.JSONParsingErrorResponse(err)
	}

	err = s.ordering.SetItemStatus(ctx, id, req.Status)
	if err != nil {
		return failures.Handle(throw.Error(err))
	}

	return handler.ResponseWithData(http.StatusNoContent, nil)
}
