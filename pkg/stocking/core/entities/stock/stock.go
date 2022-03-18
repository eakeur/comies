package stock

import (
	"gomies/pkg/sdk/types"
	"time"
)

// MovementType is an indicator pointing out if the movement is an input our output
type MovementType int

const (
	Input  MovementType = iota
	Output MovementType = iota
)

type Movement struct {
	types.Entity

	// TargetID is an identifier for the object this stocks references to
	TargetID types.External

	// Type points out if this movement is input or output
	Type MovementType

	// Date is when the object got into the stock effectively
	Date time.Time

	// Quantity is the amount being inserted or removed from this stock
	Quantity types.Quantity

	// AdditionalData is a general-purpose space to store additional data about this entry
	AdditionalData string
}

func (m Movement) Value() types.Quantity {
	if m.Type == Output {
		return m.Quantity * -1
	}

	return m.Quantity

}

func (m Movement) Validate() error {
	if m.TargetID == types.Nil {
		return ErrMustHaveTargetID
	}

	return nil
}
