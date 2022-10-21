package menu

import (
	"comies/app/core/movement"
	"comies/app/core/types"
)

type DispatchDetails struct {
	ID        types.ID
	ProductID types.ID
	Quantity  types.Quantity
}

func DispatchProduct(listIngredients IngredientsLister, saveMovement MovementWriter) func(r DispatchDetails) error {
	return func(r DispatchDetails) error {
		list, err := listIngredients(r.ProductID)
		if err != nil {
			return err
		}

		if len(list) == 0 {
			_ = saveMovement(movement.Movement{
				ProductID: r.ProductID,
				Quantity:  r.Quantity,
				Type:      movement.OutputType,
			})

			return nil
		}

		for _, ing := range list {
			_ = DispatchProduct(listIngredients, saveMovement)(DispatchDetails{
				ID:        r.ID,
				ProductID: ing.ProductID,
				Quantity:  r.Quantity * ing.Quantity,
			})
		}

		return nil
	}

}
