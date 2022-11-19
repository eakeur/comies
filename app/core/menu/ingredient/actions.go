package ingredient

import (
	"comies/app/core/types"
	"context"
)

//go:generate moq -fmt goimports -out actions_mock.go . Actions:ActionsMock

type Actions interface {
	ListByProductID(ctx context.Context, productID types.ID) ([]Ingredient, error)
	Create(ctx context.Context, i Ingredient) error
	Remove(ctx context.Context, productID types.ID, ingredientID types.ID) error
	RemoveByProductID(ctx context.Context, productID types.ID) error
}
