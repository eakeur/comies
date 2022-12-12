package movement

import (
	"comies/core/types"
	"time"
)

type Movement struct {
	ID          types.ID       `json:"id"`
	ItemID      types.ID       `json:"item_id"`
	AgentID     types.ID       `json:"agent_id"`
	Type        types.Type     `json:"type"`
	Date        time.Time      `json:"date"`
	Price       types.Currency `json:"price"`
	Description string         `json:"description"`
}

func (m Movement) WithID(id types.ID) Movement {
	m.ID = id
	return m
}

func (m Movement) WithDate(t time.Time) Movement {
	m.Date = t.UTC()
	return m
}

func (m Movement) WithDescription(description string) Movement {
	m.Description = description
	return m
}

func (m Movement) Validate() (Movement, error) {
	if m.Price <= 0 {
		return m, ErrInvalidPrice
	}
	return m, ValidateType(m.Type)
}
