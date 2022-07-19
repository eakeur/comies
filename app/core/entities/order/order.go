package order

import (
	"comies/app/sdk/types"
	"time"
)

const (
	NoStatus              Status = 0
	InTheCartStatus       Status = 10 // "IN_THE_CART"
	PendingStatus         Status = 20 // "PENDING"
	PreparingStatus       Status = 30 // "PREPARING"
	WaitingTakeoutStatus  Status = 40 // "WAITING_TAKEOUT"
	WaitingDeliveryStatus Status = 50 // "WAITING_DELIVERY"
	DeliveringStatus      Status = 60 // "DELIVERING"
	FinishedStatus        Status = 70 // "FINISHED"
	CanceledStatus        Status = 80 // "CANCELED"
)

const (
	NoDeliveryMode       DeliveryMode = 0
	TakeoutDeliveryMode  DeliveryMode = 10 // "TAKEOUT"
	DeliveryDeliveryMode DeliveryMode = 20 // "DELIVERY"
)

type (
	Status       int
	DeliveryMode int

	Order struct {
		ID             types.ID
		Identification string
		PlacedAt       time.Time
		Status         Status
		DeliveryMode   DeliveryMode
		Observations   string
		FinalPrice     types.Currency
		Address        string
		Phone          string
	}

	FlowUpdate struct {
		ID         types.ID
		OrderID    types.ID
		Status     Status
		OccurredAt time.Time
	}

	Filter struct {
		Status       []Status
		PlacedBefore time.Time
		PlacedAfter  time.Time
		DeliveryMode DeliveryMode
	}
)

func (o Order) Validate() error {
	return nil
}
