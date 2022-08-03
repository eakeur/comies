package menu

import (
	"comies/app/gateway/api/failures"
	"comies/app/gateway/api/handler"
	"comies/app/sdk/throw"
	"context"
	"net/http"
)

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
