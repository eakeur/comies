package menu

import (
	"comies/app/core/entities/ingredient"
	"comies/app/gateway/api/response"
	"comies/app/sdk/throw"
	"context"
	"net/http"
)

func (s Service) CreateIngredient(ctx context.Context, ing ingredient.Ingredient) response.Response {
	ing, err := s.menu.CreateIngredient(ctx, ing)
	if err != nil {
		return failures.Handle(throw.Error(err))
	}

	return response.WithData(http.StatusCreated, AdditionResult{ID: ing.ID})
}
