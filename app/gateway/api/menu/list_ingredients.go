package menu

import (
	"comies/app/gateway/api/gen/menu"
	"comies/app/sdk/throw"
	"comies/app/sdk/types"
	"context"
)

func (s service) ListIngredients(ctx context.Context, in *menu.ListIngredientsRequest) (*menu.ListIngredientsResponse, error) {
	list, err := s.menu.ListProductIngredients(ctx, types.ID(in.ProductID))
	if err != nil {
		return nil, throw.Error(err)
	}

	var ingredients []*menu.IngredientsListItem
	for _, p := range list {
		ingredients = append(ingredients, &menu.IngredientsListItem{
			Id:           int64(p.ID),
			ProductID:    int64(p.ProductID),
			IngredientID: int64(p.IngredientID),
			Quantity:     int64(p.Quantity),
			Optional:     p.Optional,
		})
	}

	return &menu.ListIngredientsResponse{
		Ingredients: ingredients,
	}, nil
}
