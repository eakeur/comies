package menu

import (
	"comies/app/core/id"
	"comies/app/core/types"
)

type Reservation struct {
	ID         id.ID                   `json:"id"`
	ReserveFor id.ID                   `json:"reserve_for"`
	ProductID  id.ID                   `json:"product_id"`
	Quantity   types.Quantity          `json:"quantity"`
	Specifics  IngredientSpecification `json:"specifics"`
}

type ReservationFailure struct {
	For       id.ID `json:"for"`
	ProductID id.ID `json:"product_id"`
	Error     error `json:"error"`
}

func ReservationToReservedMovement(r Reservation) Movement {
	return AssignMovementQuantity(Movement{
		ProductID: r.ProductID,
		quantity:  r.Quantity,
		AgentID:   r.ID,
		Type:      ReservedMovementType,
	}, r.Quantity)
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
