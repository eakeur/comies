package menu

import (
	"comies/app/gateway/api/handler"
	"comies/app/gateway/api/response"
	"comies/app/sdk/throw"
	"context"
	"net/http"
)

func (s Service) GetProductIngredients(ctx context.Context, params handler.RouteParams) response.Response {
	id, err, res := convertToID(params["product_id"])
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

	return response.WithData(http.StatusOK, ingredients)
}
