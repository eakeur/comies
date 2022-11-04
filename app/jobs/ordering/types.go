package ordering

import (
	"comies/app/core/ordering/item"
	"comies/app/core/types"
	"time"
)

type (
	OrderConfirmation struct {
		items           []item.Item
		DeliveryType    types.Type
		Observations    string
		CustomerName    string
		CustomerPhone   string
		CustomerAddress string
		Time            time.Time
	}
)
