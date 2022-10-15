package menu

import (
	"comies/app/core/id"
	"comies/app/core/ingredient"
	"comies/app/core/types"
	"comies/app/handler/rest"
	"comies/app/workflows/menu"
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
// @Success     201         {object} rest.Response{data=IngredientAdditionResult{}}
// @Failure     400         {object} rest.Response{error=rest.Error{}} "INVALID_ID"
// @Failure     412         {object} rest.Response{error=rest.Error{}} "INGREDIENT_INVALID_INGREDIENT_ID, INGREDIENT_INVALID_PRODUCT_ID, INGREDIENT_ZERO_QUANTITY, INGREDIENT_INVALID_PRODUCT_TYPE, INGREDIENT_INVALID_INGREDIENT_TYPE"
// @Failure     500         {object} rest.Response{error=rest.Error{}} "ERR_INTERNAL_SERVER_ERROR"
// @Router      /menu/products/{product_id}/ingredients [POST]
func CreateIngredient(ctx context.Context, r *http.Request) rest.Response {

	var i CreateIngredientRequest
	err := json.NewDecoder(r.Body).Decode(&i)
	if err != nil {
		return rest.JSONParsingErrorResponse(err)
	}

	productID, err := rest.GetResourceIDFromURL(r, "product_id")
	if err != nil {
		return rest.IDParsingErrorResponse(err)
	}

	ingredientID, err := rest.ConvertToID(i.IngredientID)
	if err != nil {
		return rest.IDParsingErrorResponse(err)
	}

	ing, err := menu.CreateIngredient(ctx, i.ToIngredient(productID, ingredientID))
	if err != nil {
		return rest.Fail(err)
	}

	return rest.ResponseWithData(http.StatusCreated, IngredientAdditionResult{ID: ing.ID.String()})
}

type CreateIngredientRequest struct {
	// IngredientID
	IngredientID string `json:"ingredient_id"`
	// Quantity
	Quantity types.Quantity `json:"quantity"`
	// Optional
	Optional bool `json:"optional"`
}

func (i *CreateIngredientRequest) ToIngredient(productID, ingredientID id.ID) ingredient.Ingredient {
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
