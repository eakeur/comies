package ingredient

import (
	"context"
	"gomies/app/sdk/types"
)

//go:generate moq -fmt goimports -out actions_mock.go . Actions:ActionsMock

type Actions interface {
	ListIngredients(ctx context.Context, productID types.ID) ([]Ingredient, error)
	SaveIngredient(ctx context.Context, ingredient Ingredient) (Ingredient, error)
	RemoveIngredient(ctx context.Context, ingredientID types.ID) error
	RemoveAllIngredients(ctx context.Context, productID types.ID) error
}
