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

const (
	PreparingItemStatus Status = 0
	DoneItemStatus      Status = 10
	FailedItemStatus    Status = 20
)

const (
	DeliveryDeliveryMode DeliveryMode = 0
	TakeoutDeliveryMode  DeliveryMode = 10
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
		ItemID       types.ID
		ProductID    types.ID
		Quantity     types.Quantity
		Status       PreparationStatus
		Observations string
		Details      ContentDetails
	}

	ContentDetails struct {
		IgnoreIngredients  []types.ID
		ReplaceIngredients []ContentSubstitution
	}

	ContentSubstitution struct {
		From types.ID
		To   types.ID
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

func (i Item) Validate() error {

	productsLen := len(i.Products)
	if productsLen == 0 {
		return ErrMissingProducts
	}

	return nil
}
