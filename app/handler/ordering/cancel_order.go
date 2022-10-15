package ordering

import (
	"comies/app/handler/rest"
	"comies/app/workflows/ordering"
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
// @Failure     400         {object} rest.Response{error=rest.Error{}} "INVALID_ID"
// @Failure     500         {object} rest.Response{error=rest.Error{}} "ERR_INTERNAL_SERVER_ERROR"
// @Router      /ordering/orders/{order_id} [DELETE]
func CancelOrder(ctx context.Context, r *http.Request) rest.Response {
	id, err := rest.GetResourceIDFromURL(r, "order_id")
	if err != nil {
		return rest.IDParsingErrorResponse(err)
	}

	err = ordering.CancelOrder(ctx, id)
	if err != nil {
		return rest.Fail(err)
	}

	return rest.ResponseWithData(http.StatusNoContent, nil)
}
