package product

import "gomies/app/sdk/types"

type (
	IngredientInput struct {
		Quantity     types.Quantity
		IngredientID types.ID
	}

	Substitutions map[types.ID]types.ID

	ApproveSaleRequest struct {
		ProductID            types.ID
		Quantity             types.Quantity
		Price                types.Currency
		IngredientsToIgnore  []types.ID
		IngredientsToReplace Substitutions
	}

	ApproveSaleResponse struct {
		RemainingStock types.Quantity
		Price          types.Quantity
	}

	IngredientToVerify struct {
		IngredientID      types.ID
		Quantity          types.Quantity
		AvailableQuantity types.Quantity
	}
)
