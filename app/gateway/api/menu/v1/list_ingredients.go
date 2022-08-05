package v1

import (
	"comies/app/gateway/api/failures"
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
// @Param       product_key path     string false "The product ID"
// @Success     200         {object} handler.Response{data=[]ListProductsResponse{}}
// @Success     200         {object} handler.Response{data=[]ListRunningOutProductsResponse{}}
// @Failure     500         {object} handler.Response{error=handler.Error{}} "ERR_INTERNAL_SERVER_ERROR: Happens if an unexpected error happens on the API side"
// @Router      /menu/products/{product_id}/ingredients [GET]
func (s Service) GetProductIngredients(ctx context.Context, r *http.Request) handler.Response {
	id, err, res := handler.GetResourceIDFromURL(r, "product_id")
	if err != nil {
		return res
	}

	list, err := s.menu.ListIngredients(ctx, id)
	if err != nil {
		return failures.Handle(throw.Error(err))
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
