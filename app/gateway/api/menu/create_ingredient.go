package menu

import (
	"comies/app/core/entities/ingredient"
	"comies/app/gateway/api/failures"
	"comies/app/gateway/api/handler"
	"comies/app/sdk/throw"
	"context"
	"encoding/json"
	"net/http"
)

func (s Service) CreateIngredient(ctx context.Context, r *http.Request) handler.Response {

	var i Ingredient
	err := json.NewDecoder(r.Body).Decode(&i)
	if err != nil {
		return handler.JSONParsingErrorResponse(err)
	}

	productID, e, res := handler.ConvertToID(i.ProductID)
	if e != nil {
		return res
	}
	ingredientID, e, res := handler.ConvertToID(i.ProductID)
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

	return handler.ResponseWithData(http.StatusCreated, AdditionResult{ID: ing.ID.String()})
}
