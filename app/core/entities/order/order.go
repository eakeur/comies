package order

import (
	"gomies/pkg/sdk/types"
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
	Status       int
	ItemStatus   int
	DeliveryMode int

	Order struct {
		types.Entity
		types.Store
		PlacedAt     time.Time
		AddressID    types.UID
		CustomerID   types.UID
		Status       Status
		DeliveryMode DeliveryMode
	}

	Item struct {
		types.Entity
		OrderID    types.ID
		OrderUID   types.UID
		Products   []Content
		ItemStatus ItemStatus
		Price      types.Currency
		FinalPrice types.Currency
		Discount   types.Discount
	}

	Content struct {
		types.Entity
		ItemID    types.ID
		ItemUID   types.UID
		ProductID types.UID
		Quantity  types.Quantity
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
