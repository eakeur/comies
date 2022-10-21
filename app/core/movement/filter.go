package movement

import (
	"comies/app/core/types"
	"time"
)

type (
	Filter struct {
		ProductID              types.ID
		InitialDate, FinalDate time.Time
	}
)
