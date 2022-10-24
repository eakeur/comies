package status

import (
	"comies/app/core/types"
	"time"
)

type Status struct {
	OrderID    types.ID
	Value      types.Status
	OccurredAt time.Time
}

func (s Status) Validate() error {
	if err := s.OrderID.Validate(); err != nil {
		return err
	}

	if err := ValidateStatus(s.Value); err != nil {
		return err
	}

	return nil
}
