package stock

import (
	"gomies/app/core/types/entity"
	"gomies/app/core/types/id"
	"gomies/app/core/types/quantity"
	"time"
)

// MovementType is an indicator pointing out if the movement is an input our output
type MovementType int

const (
	Input  MovementType = iota
	Output MovementType = iota
)

type Movement struct {
	entity.Entity

	// TargetID is an identifier for the object this stocks references to
	TargetID id.External

	// Type points out if this movement is input or output
	Type MovementType

	// Date is when the object got into the stock effectively
	Date time.Time

	// Quantity is the amount being inserted or removed from this stock
	Quantity quantity.Quantity

	// AdditionalData is a general-purpose space to store additional data about this entry
	AdditionalData string
}
