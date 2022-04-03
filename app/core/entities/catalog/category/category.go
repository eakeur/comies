package category

import "gomies/pkg/sdk/types"

type Category struct {
	types.Entity
	Name        string
	Code        string
	Description string
	types.Store
}

func (c Category) Validate() error {
	if len(c.Code) < 2 || len(c.Code) > 12 {
		return ErrInvalidCode
	}

	if len(c.Name) < 2 || len(c.Name) > 60 {
		return ErrInvalidName
	}

	return nil
}
