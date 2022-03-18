package stock

import (
	"gomies/pkg/sdk/types"
	"time"
)

const (
	InputMovement  MovementType = iota
	OutputMovement MovementType = iota
)

type (

	Movement struct {
		types.Entity

		// TargetID is an identifier for the object this stocks references to
		TargetID types.UID

		// Type points out if this movement is input or output
		Type MovementType

		// Date is when the object got into the stock effectively
		Date time.Time

		// Quantity is the amount being inserted or removed from this stock
		Quantity types.Quantity

		// AdditionalData is a general-purpose space to store additional data about this entry
		AdditionalData string
	}

)

func (m Movement) Value() types.Quantity {
	if m.Type == OutputMovement {
		return m.Quantity * -1
	}

	return m.Quantity

}

func (m Movement) Validate() error {
	if m.TargetID.Empty() {
		return ErrMissingResourceID
	}

	return nil
}


