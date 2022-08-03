package menu

import (
	"comies/app/core/entities/movement"
	"comies/app/gateway/api/failures"
	"comies/app/gateway/api/handler"
	"comies/app/sdk/throw"
	"context"
	"encoding/json"
	"net/http"
)

func (s Service) CreateMovement(ctx context.Context, r *http.Request) handler.Response {

	var mov Movement
	err := json.NewDecoder(r.Body).Decode(&mov)
	if err != nil {
		return handler.JSONParsingErrorResponse(err)
	}

	productID, e, res := handler.ConvertToID(mov.ProductID)
	if e != nil {
		return res
	}
	agentID, _, _ := handler.ConvertToID(mov.AgentID)

	bal, err := s.menu.CreateMovement(ctx, movement.Movement{
		ProductID: productID,
		Type:      mov.Type,
		Date:      mov.Date,
		Quantity:  mov.Quantity,
		PaidValue: mov.PaidValue,
		AgentID:   agentID,
	})
	if err != nil {
		return failures.Handle(throw.Error(err))
	}

	return handler.ResponseWithData(http.StatusCreated, MovementAdditionResult{ID: bal.ID.String(), Balance: bal.Count})
}
