package ordering

import (
	"comies/app/gateway/api/handler"
	"comies/app/gateway/api/response"
	"comies/app/sdk/throw"
	"context"
	"net/http"
)

func (s Service) SetItemStatus(ctx context.Context, params handler.RouteParams, req SetItemStatusRequest) response.Response {
	id, err, res := convertToID(params["item_id"])
	if err != nil {
		return res
	}

	err = s.ordering.SetItemStatus(ctx, id, req.Status)
	if err != nil {
		return failures.Handle(throw.Error(err))
	}

	return response.WithData(http.StatusNoContent, nil)
}
