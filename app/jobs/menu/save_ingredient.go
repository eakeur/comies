package menu

import (
	"comies/app/core/menu"
	"comies/app/data/ids"
	"comies/app/data/ingredients"
	"comies/app/data/products"
	"context"
)

func CreateIngredient(ctx context.Context, i menu.Ingredient) (menu.Ingredient, error) {

	i.ID = ids.Create()

	prd, err := products.GetByID(ctx, i.ProductID)
	if err != nil || !menu.IsComposite(prd.Type) {
		return menu.Ingredient{}, menu.ErrInvalidCompositeID
	}

	ing, err := products.GetByID(ctx, i.IngredientID)
	if err != nil || menu.IsOutput(ing.Type){
		return menu.Ingredient{}, menu.ErrInvalidComponentID
	}

	return i, ingredients.Create(ctx, i)
}
