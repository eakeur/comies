package ordering

import "time"

type OrderFilter struct {
	Status       []Status
	PlacedBefore time.Time
	PlacedAfter  time.Time
	DeliveryType Type
}