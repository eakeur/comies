package ordering

import (
	"comies/app/core/entities/item"
	"comies/app/core/entities/order"
	"comies/app/sdk/types"
)

type (
	Reservation struct {
		ID        types.ID
		ProductID types.ID
		Quantity  types.Quantity
		Ignore    []item.Ignoring
		Replace   []item.Replacement
		Failures  []ItemFailed
	}

	ItemFailed struct {
		ProductID types.ID
		Want      types.Quantity
		Got       types.Quantity
		Error     error
	}

	ItemAdditionResult struct {
		Item      item.Item
		Succeeded []Reservation
		Failed    []Reservation
	}

	OrderConfirmation struct {
		OrderID      types.ID
		DeliveryMode order.DeliveryMode
	}
)
