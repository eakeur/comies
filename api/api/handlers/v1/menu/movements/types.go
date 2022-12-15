package movements

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
