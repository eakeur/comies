package category

import "gomies/app/core/types/entity"

type Category struct {
	entity.Entity
	Name        string
	Code        string
	Description string
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
