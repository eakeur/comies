package menu

import (
	"comies/app/api/request"
	"comies/app/api/send"
	"comies/app/core/types"
	"comies/app/jobs/menu"
	"context"
	"net/http"
	"time"
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
func CreateMovement(ctx context.Context, r request.Request) send.Response {

	var mov CreateMovementRequest
	err := r.JSONBody(&mov)
	if err != nil {
		return send.JSONError(err)
	}

	productID, err := r.IDParam("product_id")
	if err != nil {
		return send.IDError(err)
	}

	bal, err := menu.CreateMovement(ctx, menu.Movement{
		ProductID: productID,
		AgentID:   mov.AgentID,
		Type:      mov.Type,
		Date:      mov.Date,
	}, mov.Quantity)
	if err != nil {
		return send.FromError(err)
	}

	return send.Data(http.StatusCreated, Balance{Balance: bal.Count}, send.WithHeaders(map[string]string{
		"id": bal.ID.String(),
	}))
}

type CreateMovementRequest struct {
	// Type points out if this movement is input or output
	Type menu.Type `json:"type"`

	// Date is when the object got into the stock effectively
	Date time.Time `json:"date"`

	// Quantity is the amount being inserted or removed from this stock
	Quantity types.Quantity `json:"quantity"`

	// PaidValue is how much was paid/received for this resource
	PaidValue types.Currency `json:"paid_value"`

	// AgentID is the entity from this resource came from or is going to
	AgentID types.ID `json:"agent_id"`
}

type Balance struct {
	Balance types.Quantity `json:"balance"`
}
