package order

import (
	"comies/core/types"
	"time"
)

type Filter struct {
	Status       []types.Status
	PlacedBefore time.Time
	PlacedAfter  time.Time
	DeliveryType types.Type
}
