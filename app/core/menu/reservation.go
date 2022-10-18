package menu

import (
	"comies/app/core/types"
)

type Reservation struct {
	ID         types.ID                `json:"id"`
	ReserveFor types.ID                `json:"reserve_for"`
	ProductID  types.ID                `json:"product_id"`
	Quantity   types.Quantity          `json:"quantity"`
	Specifics  IngredientSpecification `json:"specifics"`
}

type ReservationFailure struct {
	For       types.ID `json:"for"`
	ProductID types.ID `json:"product_id"`
	Error     error    `json:"error"`
}

func ReservationToReservedMovement(r Reservation) Movement {
	return Movement{
		ProductID: r.ProductID,
		quantity:  r.Quantity,
		AgentID:   r.ID,
		Type:      ReservedMovementType,
	}
}

func ReservationToIngredientReservation(r Reservation, i Ingredient) Reservation {
	return Reservation{
		ID:         r.ID,
		Specifics:  r.Specifics,
		ReserveFor: i.ProductID,
		ProductID:  i.IngredientID,
		Quantity:   i.Quantity,
	}
}
