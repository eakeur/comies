package ordering

import (
	"comies/app/gateway/api/failures"
	"comies/app/gateway/api/handler"
	"comies/app/sdk/throw"
	"context"
	"encoding/json"
	"net/http"
)

func (s Service) SetOrderDeliveryMode(ctx context.Context, r *http.Request) handler.Response {
	id, err, res := handler.GetResourceIDFromURL(r, "order_id")
	if err != nil {
		return res
	}

	var req SetOrderDeliveryModeRequest
	err = json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return handler.JSONParsingErrorResponse(err)
	}

	err = s.ordering.SetOrderDeliveryMode(ctx, id, req.Mode)
	if err != nil {
		return failures.Handle(throw.Error(err))
	}

	return handler.ResponseWithData(http.StatusNoContent, nil)
}
