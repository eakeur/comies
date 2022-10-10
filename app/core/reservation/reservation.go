package reservation

import (
	"comies/app/core/id"
	"comies/app/core/types"
)

type (
	Reservation struct {
		ID         id.ID
		ReserveFor id.ID
		ProductID  id.ID
		Quantity   types.Quantity
		Ignore     []id.ID
		Replace    map[id.ID]id.ID
	}

	Failure struct {
		For       id.ID
		ProductID id.ID
		Error     error
	}
)
