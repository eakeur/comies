package ordering

import (
	"comies/app/api/request"
	"comies/app/api/send"
	"comies/app/core/ingredient"
	"comies/app/core/types"
	"comies/app/jobs/ordering"
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
func AddToOrder(ctx context.Context, r request.Request) send.Response {

	var i AddToOrderRequest
	err := json.NewDecoder(r.Body).Decode(&i)
	if err != nil {
		return send.JSONError(err)
	}

	it, err := r.IDParam("order_id")
	if err != nil {
		send.IDError(err)
	}

	res, id, err := ordering.AddToOrder(ctx, ordering.Item{
		OrderID:       it,
		ProductID:     i.ProductID,
		Quantity:      i.Quantity,
		Observations:  i.Observations,
		Specification: i.Specifications,
	})
	if err != nil {
		return send.FromError(err)
	}

	if size := len(res); size > 0 {
		return send.Data(http.StatusPreconditionFailed, res)
	}

	defer r.Commit(ctx)
	return send.CreatedWithID(id)
}

type AddToOrderRequest struct {
	ProductID      types.ID                           `json:"product_id"`
	Quantity       types.Quantity                     `json:"quantity"`
	Observations   string                             `json:"observations"`
	Specifications ingredient.IngredientSpecification `json:"specifications"`
}
