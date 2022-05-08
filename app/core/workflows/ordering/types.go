package ordering

import (
	"gomies/app/core/entities/order"
	"gomies/app/sdk/types"
)

type (
	Reservation struct {
		ID        types.ID
		ProductID types.ID
		Quantity  types.Quantity
		Ignore    []types.ID
		Replace   map[types.ID]types.ID
		Failures  []ItemFailed
	}

	ItemFailed struct {
		ProductID types.ID
		Want      types.Quantity
		Got       types.Quantity
		Error     error
	}

	ItemAdditionResult struct {
		Item        order.Item
		Reservation []Reservation
	}

	OrderConfirmation struct {
		OrderID      types.ID
		AddressID    types.ID
		Status       order.Status
		DeliveryMode order.DeliveryMode
	}
)
