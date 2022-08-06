package menu

import (
	"comies/app/gateway/api/handler"
	"comies/app/sdk/throw"
	"context"
	"net/http"
)

// GetProductIngredients fetches all product ingredients.
//
// @Summary     Fetches ingredients
// @Description Fetches all product ingredients.
// @Tags        Product
// @Param       product_id path     string false "The product ID"
// @Success     200         {object} handler.Response{data=[]Ingredient{}}
// @Failure     500         {object} handler.Response{error=handler.Error{}} "ERR_INTERNAL_SERVER_ERROR"
// @Router      /menu/products/{product_id}/ingredients [GET]
func (s Service) GetProductIngredients(ctx context.Context, r *http.Request) handler.Response {
	id, err := handler.GetResourceIDFromURL(r, "product_id")
	if err != nil {
		return handler.IDParsingErrorResponse(err)
	}

	list, err := s.menu.ListIngredients(ctx, id)
	if err != nil {
		return handler.Fail(throw.Error(err))
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

	return handler.ResponseWithData(http.StatusOK, ingredients)
}
