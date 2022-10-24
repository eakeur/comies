package ordering

import (
	"comies/app/core/item"
	"comies/app/core/order"
	"comies/app/core/reservation"
	"comies/app/core/types"
)

type (
	OrderConfirmation struct {
		OrderID      types.ID
		DeliveryMode order.DeliveryMode
	}
)