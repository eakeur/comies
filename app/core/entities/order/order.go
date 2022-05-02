package order

import (
	"gomies/app/sdk/listing"
	"gomies/app/sdk/types"
	"time"
)

const (
	PendingStatus         Status = iota
	PreparingStatus       Status = iota
	WaitingTakeoutStatus  Status = iota
	WaitingDeliveryStatus Status = iota
	DeliveringStatus      Status = iota
	FinishedStatus        Status = iota
	CanceledStatus        Status = iota
)

const (
	PreparingItemStatus Status = iota
	DoneItemStatus      Status = iota
	FailedItemStatus    Status = iota
)

const (
	DeliveryDeliveryMode DeliveryMode = iota
	TakeoutDeliveryMode  DeliveryMode = iota
)

type (
	Status            int
	PreparationStatus int
	DeliveryMode      int

	Order struct {
		types.Entity
		types.Store
		PlacedAt     time.Time
		AddressID    types.ID
		CustomerID   types.ID
		Status       Status
		DeliveryMode DeliveryMode
	}

	Item struct {
		types.Entity
		OrderID    types.ID
		Products   []Content
		ItemStatus PreparationStatus
		Price      types.Currency
		FinalPrice types.Currency
		Discount   types.Discount
	}

	Content struct {
		types.Entity
		ItemID    types.ID
		ProductID types.ID
		Quantity  types.Quantity
		Status    PreparationStatus
	}

	Filter struct {
		Status       Status
		PlacedBehore time.Time
		PlacedAfter  time.Time
		CustomerID   types.ID
		AddressID    types.ID
		DeliverMode  DeliveryMode
		listing.Filter
	}
)

func (o Order) Validate() error {
	return nil
}

func (i Item) Validate() error {

	productsLen := len(i.Products)
	if productsLen == 0 {
		return ErrMissingProducts
	}

	return nil
}
