package ingredients

import (
	"comies/api/request"
	"comies/api/send"
	"context"
	"net/http"
)

func (h Handler) Remove(ctx context.Context, r request.Request) send.Response {

	productID, err := r.IDParam("product_id")
	if err != nil {
		return send.IDError(err)
	}

	ingredientID, err := r.IDParam("ingredient_id")
	if err != nil {
		return send.IDError(err)
	}

	err = h.ingredients.RemoveIngredient(ctx, productID, ingredientID)
	if err != nil {
		return send.FromError(err)
	}

	r.Commit(ctx)

	return send.Data(http.StatusNoContent, nil)
}
