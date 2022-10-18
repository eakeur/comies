package ordering

import (
	"comies/app/core/types"
	"time"
)

type Flow struct {
	OrderID    types.ID
	Status     Status
	OccurredAt time.Time
}

func ValidateFlow(f Flow) error {
	if err := types.ValidateID(f.OrderID); err != nil {
		return err
	}

	if err := ValidateOrderStatus(f.Status); err != nil {
		return err
	}

	return nil
}

func NewOrderFlow(o Order) Flow {
	return NewFlow(o.ID, PreparingOrderStatus)
}

func NewFlow(orderID types.ID, status Status) Flow {
	return Flow{
		OrderID:    orderID,
		Status:     status,
		OccurredAt: time.Now().UTC(),
	}
}
