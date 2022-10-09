package menu

import (
	"comies/app/gateway/api/handler"
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
// @Failure     400 {object} handler.Response{error=handler.Error{}} "INVALID_ID"
// @Failure     500 {object} handler.Response{error=handler.Error{}} "ERR_INTERNAL_SERVER_ERROR"
// @Router      /menu/products/{product_id}/movements/{id} [DELETE]
func RemoveProductMovement(ctx context.Context, r *http.Request) handler.Response {
	id, err := handler.GetResourceIDFromURL(r, "movement_id")
	if err != nil {
		return handler.IDParsingErrorResponse(err)
	}

	err = menu.RemoveMovement(ctx, id)
	if err != nil {
		return handler.Fail(err)
	}

	return handler.ResponseWithData(http.StatusNoContent, nil)
}
