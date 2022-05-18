package order

import (
	"gomies/app/sdk/listing"
	"gomies/app/sdk/types"
	"time"
)

const (
	InTheCartStatus       Status = 0
	PendingStatus         Status = 10
	PreparingStatus       Status = 20
	WaitingTakeoutStatus  Status = 30
	WaitingDeliveryStatus Status = 40
	DeliveringStatus      Status = 50
	FinishedStatus        Status = 60
	CanceledStatus        Status = 70
)

type (
	Status       int
	DeliveryMode int

	Order struct {
		types.Entity
		types.Store
		PlacedAt     time.Time
		AddressID    types.ID
		CustomerID   types.ID
		Status       Status
		DeliveryMode DeliveryMode
	}

	Filter struct {
		Status       Status
		PlacedBefore time.Time
		PlacedAfter  time.Time
		CustomerID   types.ID
		AddressID    types.ID
		DeliveryMode DeliveryMode
		listing.Filter
	}
)

func (o Order) Validate() error {
	return nil
}
