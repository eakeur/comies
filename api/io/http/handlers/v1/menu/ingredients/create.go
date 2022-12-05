package ingredients

import (
	"comies/core/menu/ingredient"
	"comies/io/http/request"
	"comies/io/http/send"
	"context"
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
func (h Handler) Create(ctx context.Context, r request.Request) send.Response {

	var i Ingredient
	err := r.JSONBody(&i)
	if err != nil {
		return send.JSONError(err)
	}

	productID, err := r.IDParam("product_id")
	if err != nil {
		return send.IDError(err)
	}

	ingredientID, err := r.IDParam("ingredient_id")
	if err != nil {
		return send.IDError(err)
	}

	_, err = h.ingredients.CreateIngredient(ctx, ingredient.Ingredient{
		ProductID:    productID,
		IngredientID: ingredientID,
		Quantity:     i.Quantity,
		Optional:     i.Optional,
	})

	if err != nil {
		return send.FromError(err)
	}

	r.Commit(ctx)

	return send.Data(http.StatusCreated, nil)
}
