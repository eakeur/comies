package menu

import (
	"comies/app/gateway/api/failures"
	"comies/app/gateway/api/handler"
	"comies/app/sdk/throw"
	"context"
	"net/http"
)

func (s Service) RemoveProductIngredient(ctx context.Context, r *http.Request) handler.Response {
	id, err, res := handler.GetResourceIDFromURL(r, "product_id")
	if err != nil {
		return res
	}

	err = s.menu.RemoveIngredient(ctx, id)
	if err != nil {
		return failures.Handle(throw.Error(err))
	}

	return handler.ResponseWithData(http.StatusNoContent, nil)
}
