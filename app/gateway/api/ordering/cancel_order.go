package ordering

import (
	"comies/app/core/throw"
	"comies/app/gateway/api/handler"
	"context"
	"net/http"
)

// CancelOrder sets an order as canceled.
//
// @Summary     Cancels order
// @Description Cancels an order that has been recently ordered
// @Tags        Ordering
// @Param       order_id path     string                  false "The order ID"
// @Success     204
// @Failure     400         {object} handler.Response{error=handler.Error{}} "INVALID_ID"
// @Failure     500         {object} handler.Response{error=handler.Error{}} "ERR_INTERNAL_SERVER_ERROR"
// @Router      /ordering/orders/{order_id} [DELETE]
func (s Service) CancelOrder(ctx context.Context, r *http.Request) handler.Response {
	id, err := handler.GetResourceIDFromURL(r, "order_id")
	if err != nil {
		return handler.IDParsingErrorResponse(err)
	}

	err = s.ordering.CancelOrder(ctx, id)
	if err != nil {
		return handler.Fail(throw.Error(err))
	}

	return handler.ResponseWithData(http.StatusNoContent, nil)
}
