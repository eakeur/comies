package menu

import (
	"comies/app/handler/rest"
	"comies/app/workflows/menu"
	"context"
	"net/http"
)

// RemoveProductMovement remove a product movement from the store's menu.
//
// @Summary     Remove movement
// @Description removes a movement from the store's menu.
// @Tags        Product
// @Param       product_id path string true "The product ID"
// @Param       id         path string true "The movement ID"
// @Success     204
// @Failure     400 {object} rest.Response{error=rest.Error{}} "INVALID_ID"
// @Failure     500 {object} rest.Response{error=rest.Error{}} "ERR_INTERNAL_SERVER_ERROR"
// @Router      /menu/products/{product_id}/movements/{id} [DELETE]
func RemoveProductMovement(ctx context.Context, r *http.Request) rest.Response {
	id, err := rest.GetResourceIDFromURL(r, "movement_id")
	if err != nil {
		return rest.IDParsingErrorResponse(err)
	}

	err = menu.RemoveMovement(ctx, id)
	if err != nil {
		return rest.Fail(err)
	}

	return rest.ResponseWithData(http.StatusNoContent, nil)
}
