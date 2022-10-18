package ordering

import (
	"comies/app/core/menu"
	"comies/app/core/types"
)

type Item struct {
	ID            types.ID                     `json:"id"`
	OrderID       types.ID                     `json:"order_id"`
	Status        Status                       `json:"status"`
	ProductID     types.ID                     `json:"product_id"`
	Quantity      types.Quantity               `json:"quantity"`
	Observations  string                       `json:"observations"`
	Specification menu.IngredientSpecification `json:"specification"`
}

func ValidateItemStatus(s Status) error {
	switch s {
	case PreparingItemStatus, DoneItemStatus, FailedItemStatus:
		return nil
	default:
		return ErrInvalidStatus
	}
}

func ValidateItem(i Item) error {
	if err := ValidateItemStatus(i.Status); err != nil {
		return err
	}

	if i.Quantity <= types.QuantityMinimum {
		return ErrInvalidQuantity
	}

	if err := types.ValidateID(i.ProductID); err != nil {
		return err
	}

	if err := types.ValidateID(i.OrderID); err != nil {
		return err
	}
	return nil
}
