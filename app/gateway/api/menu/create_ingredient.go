package menu

import (
	"comies/app/core/entities/ingredient"
	"comies/app/gateway/api/failures"
	"comies/app/gateway/api/handler"
	"comies/app/sdk/throw"
	"comies/app/sdk/types"
	"context"
	"encoding/json"
	"net/http"
)

// CreateIngredient adds an ingredient relation to the store's menu.
//
// @Summary     Creates ingredient
// @Description Adds an ingredient relation to the store's menu. THe product must be of composite type
// @Tags        Product
// @Param       product_id path     string                  false "The product ID"
// @Param       ingredient  body     CreateIngredientRequest true  "The properties to define the ingredient"
// @Success     201         {object} handler.Response{data=IngredientAdditionResult{}}
// @Failure     400         {object} handler.Response{error=handler.Error{}} "INVALID_ID"
// @Failure     412         {object} handler.Response{error=handler.Error{}} "INGREDIENT_INVALID_INGREDIENT_ID, INGREDIENT_INVALID_PRODUCT_ID, INGREDIENT_ZERO_QUANTITY, INGREDIENT_INVALID_PRODUCT_TYPE, INGREDIENT_INVALID_INGREDIENT_TYPE"
// @Failure     500         {object} handler.Response{error=handler.Error{}} "ERR_INTERNAL_SERVER_ERROR"
// @Router      /menu/products/{product_id}/ingredients [POST]
func (s Service) CreateIngredient(ctx context.Context, r *http.Request) handler.Response {

	var i CreateIngredientRequest
	err := json.NewDecoder(r.Body).Decode(&i)
	if err != nil {
		return handler.JSONParsingErrorResponse(err)
	}

	productID, err := handler.GetResourceIDFromURL(r, "product_id")
	if err != nil {
		return handler.IDParsingErrorResponse(err)
	}

	ingredientID, err := handler.ConvertToID(i.IngredientID)
	if err != nil {
		return handler.IDParsingErrorResponse(err)
	}

	ing, err := s.menu.CreateIngredient(ctx, i.ToIngredient(productID, ingredientID))
	if err != nil {
		return failures.Handle(throw.Error(err))
	}

	return handler.ResponseWithData(http.StatusCreated, IngredientAdditionResult{ID: ing.ID.String()})
}

type CreateIngredientRequest struct {
	IngredientID string         `json:"ingredient_id"`
	Quantity     types.Quantity `json:"quantity"`
	Optional     bool           `json:"optional"`
}

func (i *CreateIngredientRequest) ToIngredient(productID, ingredientID types.ID) ingredient.Ingredient {
	return ingredient.Ingredient{
		ProductID:    productID,
		IngredientID: ingredientID,
		Quantity:     i.Quantity,
		Optional:     i.Optional,
	}
}

type IngredientAdditionResult struct {
	ID string `json:"id"`
}
