package movement

import (
	"comies/app/sdk/types"
	"time"
)

const (
	NoType       Type = 0
	InputType    Type = 10
	OutputType   Type = 20
	ReservedType Type = 30
)

type Type int

type Movement struct {
	ID types.ID

	// ProductID is an identifier for the stock this movement references to
	ProductID types.ID

	// Type points out if this movement is input or output
	Type Type

	// Date is when the object got into the stock effectively
	Date time.Time

	// Quantity is the amount being inserted or removed from this stock
	Quantity types.Quantity

	// PaidValue is how much was paid/received for this resource
	PaidValue types.Currency

	// AgentID is the entity from this resource came from or is going to
	AgentID types.ID
}

func (m Movement) Value() types.Quantity {
	if m.Type == OutputType || m.Type == ReservedType {
		return m.Quantity * -1
	}

	return m.Quantity

}

func (m Movement) Validate() error {
	return nil
}
