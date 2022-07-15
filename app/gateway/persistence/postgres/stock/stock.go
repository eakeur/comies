package stock

import (
	"comies/app/core/entities/stock"

	"github.com/jackc/pgx/v4/pgxpool"
)

var _ stock.Actions = actions{}

type actions struct {
	db *pgxpool.Pool
}

func NewActions(db *pgxpool.Pool) stock.Actions {
	return actions{db: db}
}
