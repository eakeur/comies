package menu

import (
	"comies/app/core/types"
	"time"
)

type Movement struct {
	ID        types.ID  `json:"id"`
	ProductID types.ID  `json:"product_id"`
	AgentID   types.ID  `json:"agent_id"`
	Type      Type      `json:"type"`
	Date      time.Time `json:"date"`
	quantity  types.Quantity
}

func MovementQuantity(m Movement) types.Quantity {
	return m.quantity
}

func AssignMovementQuantity(m Movement, val types.Quantity) Movement {
	if (m.Type == OutputMovementType || m.Type == ReservedMovementType) && val > 0 {
		val *= -1
	}

	m.quantity = val

	return m
}

func IncrementStockQuantity(actual types.Quantity, newMovement Movement) types.Quantity {
	return actual + newMovement.quantity
}

func ValidateMovement(m Movement) error {
	return ValidateMovementType(m.Type)
}

func CanStockAfford(m Movement, actual types.Quantity, p Product) error {
	if m.Type == InputMovementType && actual > p.MaximumQuantity {
		return ErrStockAlreadyFull
	}

	if m.Type == OutputMovementType && actual < p.MinimumQuantity {
		return ErrStockNegative
	}

	return ValidateMovementType(m.Type)
}

func CheckMovementTypeCompatibility(movementType, productType Type) error {
	if (productType == InputProductType && movementType == InputMovementType) || IsComposite(productType) {
		return ErrInvalidProductType
	}

	return nil
}

func ValidateMovementType(t Type) error {
	switch t {
	case InputMovementType, OutputMovementType, ReservedMovementType:
		return nil
	default:
		return ErrInvalidType
	}
}
