package category

import "gomies/app/core/types/entity"

type Category struct {
	entity.Entity
	Name        string
	Description string
}

func (c Category) Validate() error {
	return nil
}
