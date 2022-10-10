package ordering

import (
	"comies/app/core/id"
	"comies/app/core/item"
	"comies/app/core/order"
	"comies/app/core/reservation"
)

type (
	OrderNotification struct {
		order.Order
		Items []item.Item
	}

	ReservationFailure struct {
		ProductID id.ID
		Error     error
	}

	ItemAdditionResult struct {
		Failed []reservation.Failure
	}

	OrderConfirmation struct {
		OrderID      id.ID
		DeliveryMode order.DeliveryMode
	}
)
