package menu

import (
	"comies/app/core/entities/movement"
	"comies/app/gateway/api/response"
	"comies/app/sdk/throw"
	"context"
	"net/http"
)

func (s Service) CreateMovement(ctx context.Context, mov movement.Movement) response.Response {
	bal, err := s.menu.CreateMovement(ctx, mov)
	if err != nil {
		return failures.Handle(throw.Error(err))
	}

	return response.WithData(http.StatusCreated, MovementAdditionResult{ID: bal.ID})
}
