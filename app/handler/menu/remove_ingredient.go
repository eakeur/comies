package menu

import (
	"comies/app/gateway/api/handler"
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
// @Failure     400 {object} handler.Response{error=handler.Error{}} "INVALID_ID"
// @Failure     500 {object} handler.Response{error=handler.Error{}} "ERR_INTERNAL_SERVER_ERROR"
// @Router      /menu/products/{product_id}/ingredients/{id} [DELETE]
func RemoveProductIngredient(ctx context.Context, r *http.Request) handler.Response {
	id, err := handler.GetResourceIDFromURL(r, "product_id")
	if err != nil {
		return handler.IDParsingErrorResponse(err)
	}

	err = menu.RemoveIngredient(ctx, id)
	if err != nil {
		return handler.Fail(err)
	}

	return handler.ResponseWithData(http.StatusNoContent, nil)
}
