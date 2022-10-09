package item

import (
	"comies/app/core/entities/item"

	"github.com/jackc/pgx/v4/pgxpool"
)

var _ item.Actions = actions{}

type actions struct {
	db *pgxpool.Pool
}

func NewActions(db *pgxpool.Pool) item.Actions {
	return &actions{db: db}
}
