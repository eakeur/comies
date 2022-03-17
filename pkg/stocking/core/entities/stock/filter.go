package stock

import (
	"gomies/pkg/sdk/types"
	"time"
)

type Filter struct {
	TargetID    types.External
	InitialDate time.Time
	FinalDate   time.Time
}
