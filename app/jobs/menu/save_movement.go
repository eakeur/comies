package menu

import (
	"comies/app/core/movement"
)

func SaveMovement(idgen IDGenerator, fetchProduct ProductFetcher, write MovementWriter) func(m movement.Movement) (movement.Movement, error) {
	return func(m movement.Movement) (movement.Movement, error) {
		save, err := m.WithID(idgen()).AssertQuantity().Validate()
		if err != nil {
			return movement.Movement{}, err
		}

		p, err := fetchProduct(save.ProductID)
		if err != nil {
			return movement.Movement{}, err
		}

		if (p.IsInput() && m.Type == movement.OutputType) || p.IsComposite() {
			return movement.Movement{}, movement.ErrInvalidProductType
		}

		return save, write(save)
	}

}
