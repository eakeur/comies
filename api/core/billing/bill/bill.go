package bill

import (
	"comies/core/types"
	"time"
)

type Bill struct {
	ID          types.ID
	Date        time.Time
	Name        types.Text
	ReferenceID types.ID
}

func (b Bill) WithID(id types.ID) Bill {
	b.ID = id
	return b
}

func (b Bill) WithDate(t time.Time) Bill {
	b.Date = t.UTC()
	return b
}

func (b Bill) WithName(name types.Text) Bill {
	b.Name = name
	return b
}

func (b Bill) Validate() (Bill, error) {
	if len(b.Name) < 3 {
		return b, ErrInvalidName
	}

	if b.Date.IsZero() {
		return b, types.ErrInvalidDateOrPeriod
	}

	return b, nil
}
