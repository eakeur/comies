package movement

import (
	"gomies/app/sdk/types"
	"time"
)

const (
	InputMovement    Type = iota
	OutputMovement   Type = iota
	ReservedMovement Type = iota
)

type Type int

type Movement struct {
	types.Entity

	// ResourceID is an identifier for the stock this movement references to
	ResourceID types.ID

	// Type points out if this movement is input or output
	Type Type

	// Date is when the object got into the stock effectively
	Date time.Time

	// Quantity is the amount being inserted or removed from this stock
	Quantity types.Quantity

	// PaidValue is how much was paid/received for this resource
	PaidValue types.Currency

	// Agent is the entity from this resource came from or is going to
	Agent types.ID
}

func (m Movement) Value() types.Quantity {
	if m.Type == OutputMovement || m.Type == ReservedMovement {
		return m.Quantity * -1
	}

	return m.Quantity

}

func (m Movement) Validate() error {
	return nil
}
