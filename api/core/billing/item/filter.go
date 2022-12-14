package item

import (
	"comies/core/types"
)

type Filter struct {
	BillID, ReferenceID types.ID
}

func (f Filter) Validate() error {
	errBillID := f.BillID.Validate()
	errRefID := f.ReferenceID.Validate()

	if errBillID != nil && errRefID != nil {
		return errBillID
	}

	return nil
}
