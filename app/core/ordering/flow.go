package ordering

import (
	"comies/app/core/id"
	"time"
)

type Flow struct {
	OrderID    id.ID
	Status     Status
	OccurredAt time.Time
}

func ValidateFlow(f Flow) error {
	if err := id.ValidateID(f.OrderID); err != nil {
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

func NewFlow(orderID id.ID, status Status) Flow {
	return Flow{
		OrderID:    orderID,
		Status:     status,
		OccurredAt: time.Now().UTC(),
	}
}
