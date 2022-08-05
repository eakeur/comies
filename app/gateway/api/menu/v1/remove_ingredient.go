package v1

import (
	"comies/app/gateway/api/failures"
	"comies/app/gateway/api/handler"
	"comies/app/sdk/throw"
	"context"
	"net/http"
)

// RemoveProductIngredient remove a product ingredient from the store's menu.
//
// @Summary     Remove ingredient
// @Description removes an ingredient from the store's menu.
// @Tags        Product
// @Param       product_id path string true "The product ID"
// @Param       id         path string true "The ingredient ID"
// @Success     204
// @Failure     400 {object} handler.Response{error=handler.Error{}} "INVALID_ID: Happens if the product id or ingredient provided is not a valid one"
// @Failure     500 {object} handler.Response{error=handler.Error{}} "ERR_INTERNAL_SERVER_ERROR: Happens if an unexpected error happens on the API side"
// @Router      /menu/products/{product_id}/ingredients/{id} [DELETE]
func (s Service) RemoveProductIngredient(ctx context.Context, r *http.Request) handler.Response {
	id, err, res := handler.GetResourceIDFromURL(r, "ingredient_id")
	if err != nil {
		return res
	}

	err = s.menu.RemoveIngredient(ctx, id)
	if err != nil {
		return failures.Handle(throw.Error(err))
	}

	return handler.ResponseWithData(http.StatusNoContent, nil)
}
