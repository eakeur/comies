package item

import "gomies/app/sdk/types"

const (
	PreparingStatus Status = 0
	DoneStatus      Status = 10
	FailedStatus    Status = 20
)

type (
	Status int

	Item struct {
		ID         types.ID
		History    types.History
		Active     bool
		OrderID    types.ID
		Status     Status
		Price      types.Currency
		FinalPrice types.Currency
		Discount   types.Discount
	}
)
