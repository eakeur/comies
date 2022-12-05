package movements

import (
	"comies/core/menu/movement"
	"comies/io/http/request"
	"comies/io/http/send"
	"context"
)

// CreateMovement adds a movement to the store's stock.
//
// @Summary     Creates movement
// @Description Adds a movement to the store's stock
// @Tags        Product
// @Param       product_id path     string                false "The product ID"
// @Param       movement    body     CreateMovementRequest true  "The properties to define the movement"
// @Success     201         {object} rest.Response{data=MovementAdditionResult{}}
// @Failure     400         {object} rest.Response{error=rest.Error{}} "INVALID_ID"
// @Failure     412         {object} rest.Response{error=rest.Error{}} "MOVEMENT_INVALID_PRODUCT_TYPE, PRODUCT_STOCK_EMPTY, PRODUCT_STOCK_FULL"
// @Failure     500         {object} rest.Response{error=rest.Error{}} "ERR_INTERNAL_SERVER_ERROR"
// @Router      /menu/products/{product_id}/movements [POST]
func (h Handler) Create(ctx context.Context, r request.Request) send.Response {

	productID, err := r.IDParam("product_id")
	if err != nil {
		return send.IDError(err)
	}

	var m Movement
	err = r.JSONBody(&m)
	if err != nil {
		return send.JSONError(err)
	}

	id, err := h.movements.CreateMovement(ctx, movement.Movement{
		ProductID: productID,
		AgentID:   m.AgentID,
		Type:      m.Type,
		Date:      m.Date,
		Quantity:  m.Quantity,
	})
	if err != nil {
		return send.FromError(err)
	}

	r.Commit(ctx)

	return send.CreatedWithID(id)
}
