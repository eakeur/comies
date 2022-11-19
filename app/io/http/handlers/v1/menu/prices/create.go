package prices

import (
	"comies/app/core/types"
	"comies/app/io/http/request"
	"comies/app/io/http/send"
	"context"
	"net/http"
	"strconv"
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

	productID, err := r.IDParam("product_id")
	if err != nil {
		return send.IDError(err)
	}

	valueStr := r.Param("value")
	if err != nil {
		return send.IDError(err)
	}

	value, _ := strconv.Atoi(valueStr)

	err = h.prices.UpdateProductPrice(ctx, productID, types.Currency(value))

	if err != nil {
		return send.FromError(err)
	}

	r.Commit(ctx)

	return send.Data(http.StatusCreated, nil)
}
