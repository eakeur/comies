package order

import (
	"gomies/app/sdk/types"
	"time"
)

const (
	InTheCartStatus       Status = "IN_THE_CART"
	PendingStatus         Status = "PENDING"
	PreparingStatus       Status = "PREPARING"
	WaitingTakeoutStatus  Status = "WAITING_TAKEOUT"
	WaitingDeliveryStatus Status = "WAITING_DELIVERY"
	DeliveringStatus      Status = "DELIVERING"
	FinishedStatus        Status = "FINISHED"
	CanceledStatus        Status = "CANCELED"
)

const (
	TakeoutDeliveryMode  DeliveryMode = "TAKEOUT"
	DeliveryDeliveryMode DeliveryMode = "DELIVERY"
)

type (
	Status       string
	DeliveryMode string

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
		types.History
	}

	FlowUpdate struct {
		ID         types.ID
		OrderID    types.ID
		Status     Status
		OccurredAt time.Time
		History    types.History
	}

	Filter struct {
		Status       Status
		PlacedBefore time.Time
		PlacedAfter  time.Time
		DeliveryMode DeliveryMode
	}
)

func (o Order) Validate() error {
	return nil
}
