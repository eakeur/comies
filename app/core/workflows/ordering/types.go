package ordering

import (
	"gomies/app/core/entities/item"
	"gomies/app/core/entities/order"
	"gomies/app/sdk/types"
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
		Status       order.Status
		DeliveryMode order.DeliveryMode
	}
)
