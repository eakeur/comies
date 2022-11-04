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

func (o Status) WithValue(v types.Status) Status {
	o.Value = v
	return o
}

func (o Status) WithOccurredAt(d time.Time) Status {
	o.OccurredAt = d.UTC()
	return o
}

func (s Status) Validate() (Status, error) {
	if err := s.OrderID.Validate(); err != nil {
		return s, err
	}

	if err := ValidateStatus(s.Value); err != nil {
		return s, err
	}

	return s, nil
}
