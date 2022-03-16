package stock

import (
	"gomies/app/core/types/id"
	"time"
)

type Filter struct {
	TargetID    id.External
	InitialDate time.Time
	FinalDate   time.Time
}
