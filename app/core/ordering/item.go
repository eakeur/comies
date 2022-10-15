package ordering

import (
	"comies/app/core/id"
	"comies/app/core/types"
)

type Item struct {
	ID                 id.ID
	OrderID            id.ID
	Status             Status
	Price              types.Currency
	ProductID          id.ID
	Quantity           types.Quantity
	Observations       string
	Specification map[id.ID]struct {
		ChangeType Type  `json:"change_type"`
		ReplaceBy  id.ID `json:"replace_by"`
	}
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

	if err := id.ValidateID(i.ProductID); err != nil {
		return err
	}

	if err := id.ValidateID(i.OrderID); err != nil {
		return err
	}
	return nil
}
