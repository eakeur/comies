package ordering

import (
	"comies/app/gateway/api/handler"
	"comies/app/sdk/throw"
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
// @Success     201         {object} handler.Response{data=ItemAdditionResponse{}}
// @Failure     400         {object} handler.Response{error=handler.Error{}} "INVALID_ID"
// @Failure     412         {object} handler.Response{data=[]Failure{}} "Returns a list with the offending items"
// @Failure     500         {object} handler.Response{error=handler.Error{}} "ERR_INTERNAL_SERVER_ERROR"
// @Router      /ordering/orders/{order_id}/items [POST]
func (s Service) AddToOrder(ctx context.Context, r *http.Request) handler.Response {

	var i ItemAdditionRequest
	err := json.NewDecoder(r.Body).Decode(&i)
	if err != nil {
		return handler.JSONParsingErrorResponse(err)
	}

	it, err := i.ToItem(handler.GetURLParam(r, "order_id"))
	if err != nil {
		handler.IDParsingErrorResponse(err)
	}

	res, err := s.ordering.AddToOrder(ctx, it)
	if err != nil {
		return handler.Fail(throw.Error(err))
	}

	if size := len(res.Failed); size > 0 {
		return handler.ResponseWithData(http.StatusPreconditionFailed, NewFailureList(res.Failed))
	}

	return handler.ResponseWithData(http.StatusCreated, ItemAdditionResponse{ID: it.ID.String()})
}
