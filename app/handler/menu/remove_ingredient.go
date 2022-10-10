package menu

import (
	"comies/app/core/workflows/menu"
	"comies/app/handler/rest"
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
// @Failure     400 {object} rest.Response{error=rest.Error{}} "INVALID_ID"
// @Failure     500 {object} rest.Response{error=rest.Error{}} "ERR_INTERNAL_SERVER_ERROR"
// @Router      /menu/products/{product_id}/ingredients/{id} [DELETE]
func RemoveProductIngredient(ctx context.Context, r *http.Request) rest.Response {
	id, err := rest.GetResourceIDFromURL(r, "product_id")
	if err != nil {
		return rest.IDParsingErrorResponse(err)
	}

	err = menu.RemoveIngredient(ctx, id)
	if err != nil {
		return rest.Fail(err)
	}

	return rest.ResponseWithData(http.StatusNoContent, nil)
}
