package menu

import (
	"comies/api/request"
	"comies/api/send"
	"context"
	"net/http"
)

func (h Handler) RemoveItemIngredient(ctx context.Context, r request.Request) send.Response {

	productID, err := r.IDParam(ItemIDParam)
	if err != nil {
		return send.IDError(err)
	}

	ingredientID, err := r.IDParam(IngredientIDParam)
	if err != nil {
		return send.IDError(err)
	}

	err = h.menu.RemoveIngredient(ctx, productID, ingredientID)
	if err != nil {
		return send.FromError(err)
	}

	r.Commit(ctx)

	return send.Data(http.StatusNoContent, nil)
}
