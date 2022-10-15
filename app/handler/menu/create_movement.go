package menu

import (
	"comies/app/core/id"
	"comies/app/core/movement"
	"comies/app/core/types"
	"comies/app/handler/rest"
	"comies/app/workflows/menu"
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
// @Success     201         {object} rest.Response{data=MovementAdditionResult{}}
// @Failure     400         {object} rest.Response{error=rest.Error{}} "INVALID_ID"
// @Failure     412         {object} rest.Response{error=rest.Error{}} "MOVEMENT_INVALID_PRODUCT_TYPE, PRODUCT_STOCK_EMPTY, PRODUCT_STOCK_FULL"
// @Failure     500         {object} rest.Response{error=rest.Error{}} "ERR_INTERNAL_SERVER_ERROR"
// @Router      /menu/products/{product_id}/movements [POST]
func CreateMovement(ctx context.Context, r *http.Request) rest.Response {

	var mov CreateMovementRequest
	err := json.NewDecoder(r.Body).Decode(&mov)
	if err != nil {
		return rest.JSONParsingErrorResponse(err)
	}

	productID, err := rest.GetResourceIDFromURL(r, "product_id")
	if err != nil {
		return rest.IDParsingErrorResponse(err)
	}

	agentID, _ := rest.ConvertToID(mov.AgentID)

	bal, err := menu.CreateMovement(ctx, mov.ToMovement(productID, agentID))
	if err != nil {
		return rest.Fail(err)
	}

	return rest.ResponseWithData(http.StatusCreated, MovementAdditionResult{ID: bal.ID.String(), Balance: bal.Count})
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

func (m CreateMovementRequest) ToMovement(productID, agentID id.ID) movement.Movement {
	return movement.Movement{
		ProductID: productID,
		AgentID:   agentID,
		Type:      m.Type,
		Date:      m.Date,
		Quantity:  m.Quantity,
		PaidValue: m.PaidValue,
	}
}
