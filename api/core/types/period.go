package types

import "time"

type Period struct {
	Start, End time.Time
}

func (p Period) Validate() error {
	if p.End.Before(p.Start) {
		return ErrInvalidDateOrPeriod
	}

	return nil
}

func (p Period) ValidateStrict() error {
	if p.Start.IsZero() || p.End.IsZero() {
		return ErrInvalidDateOrPeriod
	}

	return p.Validate()
}
