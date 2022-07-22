package menu

import (
	"comies/app/gateway/api/handler"
	"comies/app/gateway/api/response"
	"comies/app/sdk/throw"
	"context"
	"net/http"
)

func (s Service) RemoveIngredient(ctx context.Context, params handler.RouteParams) response.Response {
	id, err, res := convertToID(params["ingredient_id"])
	if err != nil {
		return res
	}

	err = s.menu.RemoveIngredient(ctx, id)
	if err != nil {
		return failures.Handle(throw.Error(err))
	}

	return response.WithData(http.StatusNoContent, nil)
}
