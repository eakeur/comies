package menu

import (
	"comies/app/core/throw"
	"comies/app/gateway/api/handler"
	"context"
	"net/http"
)

// RemoveProduct remove a product from the store's menu.
//
// @Summary     Remove product
// @Description removes a product from the store's menu.
// @Tags        Product
// @Param       product_id path string true "The product ID"
// @Success     204
// @Failure     400 {object} handler.Response{error=handler.Error{}} "INVALID_ID"
// @Failure     500 {object} handler.Response{error=handler.Error{}} "ERR_INTERNAL_SERVER_ERROR"
// @Router      /menu/products/{product_id} [DELETE]
func (s Service) RemoveProduct(ctx context.Context, r *http.Request) handler.Response {
	id, err := handler.GetResourceIDFromURL(r, "product_id")
	if err != nil {
		return handler.IDParsingErrorResponse(err)
	}

	err = s.menu.RemoveProduct(ctx, id)
	if err != nil {
		return handler.Fail(throw.Error(err))
	}

	return handler.ResponseWithData(http.StatusNoContent, nil)
}
