package menu

import (
	"comies/app/core/entities/ingredient"
	"comies/app/gateway/api/response"
	"comies/app/sdk/throw"
	"context"
	"net/http"
)

func (s Service) CreateIngredient(ctx context.Context, i Ingredient) response.Response {
	productID, e, res := convertToID(i.ProductID)
	if e != nil {
		return res
	}
	ingredientID, e, res := convertToID(i.IngredientID)
	if e != nil {
		return res
	}

	ing, err := s.menu.CreateIngredient(ctx, ingredient.Ingredient{
		ProductID:    productID,
		IngredientID: ingredientID,
		Quantity:     i.Quantity,
		Optional:     i.Optional,
	})
	if err != nil {
		return failures.Handle(throw.Error(err))
	}

	return response.WithData(http.StatusCreated, AdditionResult{ID: ing.ID.String()})
}
