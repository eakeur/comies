package menu

import (
	"comies/app/core/entities/movement"
	"comies/app/gateway/api/response"
	"comies/app/sdk/throw"
	"context"
	"net/http"
)

func (s Service) CreateMovement(ctx context.Context, mov Movement) response.Response {
	productID, e, res := convertToID(mov.ProductID)
	if e != nil {
		return res
	}
	agentID, e, res := convertToID(mov.AgentID)
	if e != nil {
		return res
	}

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

	return response.WithData(http.StatusCreated, MovementAdditionResult{ID: bal.ID.String()})
}
