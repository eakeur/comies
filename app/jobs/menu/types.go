package menu

import (
	"comies/app/core/ingredient"
	"comies/app/core/movement"
	"comies/app/core/product"
	"comies/app/core/types"
)

type (
	IDGenerator       func() types.ID
	IngredientsLister func(types.ID) ([]ingredient.Ingredient, error)
	ProductFetcher    func(types.ID) (product.Product, error)
	ProductWriter     func(product.Product) error
	MovementWriter    func(movement.Movement) error
)

type (
	ProductSaver func(product.Product) (product.Product, error)
)
