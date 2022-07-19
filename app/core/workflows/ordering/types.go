package ordering

import (
	"comies/app/core/entities/item"
	"comies/app/core/entities/order"
	"comies/app/core/entities/reservation"
	"comies/app/sdk/types"
)

type (
	OrderNotification struct {
		order.Order
		Items []item.Item
	}

	ReservationFailure struct {
		ProductID types.ID
		Error     error
	}

	ItemAdditionResult struct {
		Failed []reservation.Failure
	}

	OrderConfirmation struct {
		OrderID      types.ID
		DeliveryMode order.DeliveryMode
	}
)
