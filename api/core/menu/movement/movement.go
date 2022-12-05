package movement

import (
	"comies/core/types"
	"time"
)

type Movement struct {
	ID        types.ID       `json:"id"`
	ProductID types.ID       `json:"product_id"`
	AgentID   types.ID       `json:"agent_id"`
	Type      types.Type     `json:"type"`
	Date      time.Time      `json:"date"`
	Quantity  types.Quantity `json:"quantity"`
	Price     types.Currency `json:"price"`
}

func (m Movement) WithID(id types.ID) Movement {
	m.ID = id
	return m
}

func (m Movement) WithDate(t time.Time) Movement {
	m.Date = t.UTC()
	return m
}

func (m Movement) AssertQuantity() Movement {
	val := m.Quantity
	if m.Type == OutputType && val > 0 {
		val *= -1
	}

	m.Quantity = val

	return m
}

func (m Movement) Validate() (Movement, error) {
	if m.Price <= 0 {
		return m, ErrInvalidPrice
	}
	return m, ValidateType(m.Type)
}
