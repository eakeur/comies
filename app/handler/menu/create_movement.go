package menu

import (
	"comies/app/core/entities/movement"
	"comies/app/core/types"
	"comies/app/gateway/api/handler"
	"context"
	"encoding/json"
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
// @Success     201         {object} handler.Response{data=MovementAdditionResult{}}
// @Failure     400         {object} handler.Response{error=handler.Error{}} "INVALID_ID"
// @Failure     412         {object} handler.Response{error=handler.Error{}} "MOVEMENT_INVALID_PRODUCT_TYPE, PRODUCT_STOCK_EMPTY, PRODUCT_STOCK_FULL"
// @Failure     500         {object} handler.Response{error=handler.Error{}} "ERR_INTERNAL_SERVER_ERROR"
// @Router      /menu/products/{product_id}/movements [POST]
func CreateMovement(ctx context.Context, r *http.Request) handler.Response {

	var mov CreateMovementRequest
	err := json.NewDecoder(r.Body).Decode(&mov)
	if err != nil {
		return handler.JSONParsingErrorResponse(err)
	}

	productID, err := handler.GetResourceIDFromURL(r, "product_id")
	if err != nil {
		return handler.IDParsingErrorResponse(err)
	}

	agentID, _ := handler.ConvertToID(mov.AgentID)

	bal, err := menu.CreateMovement(ctx, mov.ToMovement(productID, agentID))
	if err != nil {
		return handler.Fail(err)
	}

	return handler.ResponseWithData(http.StatusCreated, MovementAdditionResult{ID: bal.ID.String(), Balance: bal.Count})
}

type CreateMovementRequest struct {
	// Type points out if this movement is input or output
	Type movement.Type `json:"type"`

	// Date is when the object got into the stock effectively
	Date time.Time `json:"date"`

	// Quantity is the amount being inserted or removed from this stock
	Quantity types.Quantity `json:"quantity"`

	// PaidValue is how much was paid/received for this resource
	PaidValue types.Currency `json:"paid_value"`

	// AgentID is the entity from this resource came from or is going to
	AgentID string `json:"agent_id"`
}

type MovementAdditionResult struct {
	ID      string         `json:"id"`
	Balance types.Quantity `json:"balance"`
}

func (m CreateMovementRequest) ToMovement(productID, agentID types.ID) movement.Movement {
	return movement.Movement{
		ProductID: productID,
		AgentID:   agentID,
		Type:      m.Type,
		Date:      m.Date,
		Quantity:  m.Quantity,
		PaidValue: m.PaidValue,
	}
}
