package menu

import (
	"comies/app/core/workflows/menu"
	"comies/app/handler/rest"
	"context"
	"net/http"
)

// GetProductIngredients fetches all product ingredients.
//
// @Summary     Fetches ingredients
// @Description Fetches all product ingredients.
// @Tags        Product
// @Param       product_id path     string false "The product ID"
// @Success     200         {object} rest.Response{data=[]Ingredient{}}
// @Failure     500         {object} rest.Response{error=rest.Error{}} "ERR_INTERNAL_SERVER_ERROR"
// @Router      /menu/products/{product_id}/ingredients [GET]
func GetProductIngredients(ctx context.Context, r *http.Request) rest.Response {
	id, err := rest.GetResourceIDFromURL(r, "product_id")
	if err != nil {
		return rest.IDParsingErrorResponse(err)
	}

	list, err := menu.ListIngredients(ctx, id)
	if err != nil {
		return rest.Fail(err)
	}

	ingredients := make([]Ingredient, len(list))
	for i, p := range list {
		ingredients[i] = Ingredient{
			ID:           p.ID.String(),
			ProductID:    p.ProductID.String(),
			IngredientID: p.IngredientID.String(),
			Quantity:     p.Quantity,
			Optional:     p.Optional,
		}
	}

	return rest.ResponseWithData(http.StatusOK, ingredients)
}
