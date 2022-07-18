package menu

import (
	"comies/app/core/entities/ingredient"
	"comies/app/gateway/api/gen/menu"
	"comies/app/sdk/throw"
	"comies/app/sdk/types"
	"context"
)

func (s service) CreateIngredient(ctx context.Context, in *menu.CreateIngredientRequest) (*menu.CreateIngredientResponse, error) {
	ing, err := s.menu.AddProductIngredient(ctx, ingredient.Ingredient{
		ProductID:    types.ID(in.ProductID),
		IngredientID: types.ID(in.IngredientID),
		Quantity:     types.Quantity(in.Quantity),
		Optional:     in.Optional,
	})
	if err != nil {
		return nil, failures.HandleError(throw.Error(err))
	}

	return &menu.CreateIngredientResponse{
		Id: int64(ing.ID),
	}, nil
}
