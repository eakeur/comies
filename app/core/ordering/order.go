package ordering

import (
	"comies/app/core/id"
	"time"
)

var UndoneOrderStatuses = []Status{
	PendingOrderStatus,
	PreparingOrderStatus,
	WaitingTakeoutOrderStatus,
	WaitingDeliveryOrderStatus,
	DeliveringOrderStatus,
}

type Customer struct {
	Name    string `json:"name"`
	Address string `json:"address"`
	Phone   string `json:"phone"`
}

type Order struct {
	ID           id.ID          `json:"id"`
	PlacedAt     time.Time      `json:"placed_at"`
	Status       Status         `json:"status"`
	DeliveryType Type           `json:"delivery_mode"`
	Observations string         `json:"observations"`
	Customer     Customer       `json:"customer"`
}

func ValidateOrderStatus(s Status) error {
	switch s {
	case InTheCartOrderStatus,
		PendingOrderStatus,
		PreparingOrderStatus,
		WaitingTakeoutOrderStatus,
		WaitingDeliveryOrderStatus,
		DeliveringOrderStatus,
		FinishedOrderStatus,
		CanceledOrderStatus:
		return nil
	default:
		return ErrInvalidStatus
	}
}

func ValidateDeliveryType(t Type) error {
	switch t {
	case TakeoutDeliveryType, DeliverDeliveryType:
		return nil
	default:
		return ErrInvalidDeliveryType
	}
}

func ValidateOrder(o Order) error {
	if o.PlacedAt.IsZero() || o.PlacedAt.After(time.Now().UTC()){
		return ErrInvalidPlacementDate
	}

	if err := ValidateOrderStatus(o.Status); err != nil {
		return err
	}

	if err := ValidateDeliveryType(o.DeliveryType); err != nil {
		return err
	}

	if len(o.Customer.Name) <= 0 {
		return ErrInvalidCustomerName
	}

	if o.DeliveryType == DeliverDeliveryType {
		if len(o.Customer.Phone) <= 0 {
			return ErrInvalidCustomerPhone
		}
		if len(o.Customer.Address) <= 0 {
			return ErrInvalidCustomerAddress
		}
	}
	return nil
}

func CheckIfOrderIsPlaceable(o Order) error {
	if err := ValidateOrder(o); err != nil {
		return err
	}

	if o.Status >= PreparingOrderStatus {
		return ErrAlreadyOrdered
	}

	return nil
}

func CheckIfOrderIsCancelable(o Order) error {
	if o.Status >= PreparingOrderStatus {
		return ErrAlreadyPreparing
	}

	return nil
}

