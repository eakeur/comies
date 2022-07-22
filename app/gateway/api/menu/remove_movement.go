package menu

import (
	"comies/app/gateway/api/handler"
	"comies/app/gateway/api/response"
	"comies/app/sdk/throw"
	"context"
	"net/http"
)

func (s Service) RemoveMovement(ctx context.Context, params handler.RouteParams) response.Response {
	id, err, res := convertToID(params["movement_id"])
	if err != nil {
		return res
	}

	err = s.menu.RemoveMovement(ctx, id)
	if err != nil {
		return failures.Handle(throw.Error(err))
	}

	return response.WithData(http.StatusNoContent, nil)
}
