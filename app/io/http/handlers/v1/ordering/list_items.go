package ordering

import (
	"comies/app/api/request"
	"comies/app/api/send"
	"comies/app/data/items"
	"context"
	"net/http"
)

// ListItems
//
// @Tags        Ordering
// @Param       order_id path     string                  false "The order ID"
// @Success     200         {object} rest.Response{data=[]Item{}}
// @Failure     400         {object} rest.Response{error=rest.Error{}} "INVALID_ID"
// @Failure     500         {object} rest.Response{error=rest.Error{}} "ERR_INTERNAL_SERVER_ERROR"
// @Router      /ordering/orders/{order_id}/items [GET]
func ListItems(ctx context.Context, r request.Request) send.Response {
	id, err := r.IDParam("order_id")
	if err != nil {
		return send.IDError(err)
	}

	list, err := items.List(ctx, id)
	if err != nil {
		return send.FromError(err)
	}

	return send.Data(http.StatusOK, list)

}
