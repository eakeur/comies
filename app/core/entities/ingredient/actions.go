package ingredient

import (
	"context"
	"gomies/app/sdk/types"
)

//go:generate moq -fmt goimports -out actions_mock.go . Actions:ActionsMock

type Actions interface {
	List(ctx context.Context, productID types.ID) ([]Ingredient, error)
	Create(ctx context.Context, ingredient Ingredient) (Ingredient, error)
	Remove(ctx context.Context, ingredientID types.ID) error
	RemoveAll(ctx context.Context, productID types.ID) error
}
