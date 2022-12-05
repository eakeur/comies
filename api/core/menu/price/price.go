package price

import (
	"comies/core/types"
	"time"
)

type Price struct {
	ID       types.ID
	TargetID types.ID
	Date     time.Time
	Value    types.Currency
}

func (p Price) WithID(id types.ID) Price {
	p.ID = id
	return p
}

func (p Price) WithDate(d time.Time) Price {
	p.Date = d.UTC()
	return p
}

func (p Price) WithValue(v types.Currency) Price {
	p.Value = v
	return p
}

func (p Price) Validate() (Price, error) {
	if p.Date.IsZero() {
		return p, ErrInvalidDate
	}

	if err := p.TargetID.Validate(); err != nil {
		return p, err
	}

	if p.Value < 0 {
		return p, ErrInvalidValue
	}

	return p, nil
}
