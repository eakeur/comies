package bill

import (
	"comies/core/types"
)

type Filter struct {
	ReferenceID types.ID
	Period      types.Period
}

func (f Filter) Validate() error {
	if err := f.ReferenceID.Validate(); err != nil {
		return err
	}

	return f.Period.Validate()
}
