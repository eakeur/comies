package stock

import (
	"gomies/app/sdk/types"
	"time"
)

const (
	InputMovement    MovementType = iota
	OutputMovement   MovementType = iota
	ReservedMovement MovementType = iota
)

type (
	Movement struct {
		types.Entity

		// ResourceID is an identifier for the stock this movement references to
		ResourceID types.ID

		// Type points out if this movement is input or output
		Type MovementType

		// Date is when the object got into the stock effectively
		Date time.Time

		// Quantity is the amount being inserted or removed from this stock
		Quantity types.Quantity

		// PaidValue is how much was paid/received for this resource
		PaidValue types.Currency

		// Agent is the entity from this resource came from or is going to
		Agent types.ID

		// Batch references to a group of resources that came together
		Batch string

		// ShelfLife is the date when the resource is not usable anymore
		ShelfLife time.Time

		// AdditionalData is a general-purpose space to store additional data about this entry
		AdditionalData string
	}

	Stock struct {
		ID      types.ID
		Active  bool
		History types.History
		// TargetID is an identifier for the object this stocks references to
		TargetID types.ID
		// MaximumQuantity is how many unities of this resource the stock can support
		MaximumQuantity types.Quantity
		// MinimumQuantity is the lowest quantity of this resource the stock can have
		MinimumQuantity types.Quantity
		// Location is a brief description of where this stock is located
		Location string
	}
)

func (m Movement) Value() types.Quantity {
	if m.Type == OutputMovement || m.Type == ReservedMovement {
		return m.Quantity * -1
	}

	return m.Quantity

}

func (m Movement) Validate() error {
	return nil
}
