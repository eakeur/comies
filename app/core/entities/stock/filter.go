package stock

import (
	"gomies/app/core/types/id"
	"time"
)

type SortField int
const (
	Code         SortField = iota
	Name         SortField = iota
)

type Filter struct {
	TargetID    id.External
	InitialDate time.Time
	FinalDate   time.Time
}

