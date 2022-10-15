package ordering

import (
	"comies/app/handler/rest"
	"comies/app/workflows/ordering"
	"context"
	"encoding/json"
	"net/http"
)

// AddToOrder adds an item to the specified order.
//
// @Summary     Adds an item
// @Description Adds an item to the specified order
// @Tags        Ordering
// @Param       order_id path     string                  false "The order ID"
// @Param       item  body     ItemAdditionRequest true  "The properties defining the item"
// @Success     201         {object} rest.Response{data=ItemAdditionResponse{}}
// @Failure     400         {object} rest.Response{error=rest.Error{}} "INVALID_ID"
// @Failure     412         {object} rest.Response{data=[]Failure{}} "Returns a list with the offending items"
// @Failure     500         {object} rest.Response{error=rest.Error{}} "ERR_INTERNAL_SERVER_ERROR"
// @Router      /ordering/orders/{order_id}/items [POST]
func AddToOrder(ctx context.Context, r *http.Request) rest.Response {

	var i ItemAdditionRequest
	err := json.NewDecoder(r.Body).Decode(&i)
	if err != nil {
		return rest.JSONParsingErrorResponse(err)
	}

	it, err := i.ToItem(rest.GetURLParam(r, "order_id"))
	if err != nil {
		rest.IDParsingErrorResponse(err)
	}

	res, err := ordering.AddToOrder(ctx, it)
	if err != nil {
		return rest.Fail(err)
	}

	if size := len(res.Failed); size > 0 {
		return rest.ResponseWithData(http.StatusPreconditionFailed, NewFailureList(res.Failed))
	}

	return rest.ResponseWithData(http.StatusCreated, ItemAdditionResponse{ID: it.ID.String()})
}
