package order

import (
	"github.com/jackc/pgx/v4/pgxpool"
	"gomies/app/core/entities/order"
)

var _ order.Actions = actions{}

type actions struct {
	db *pgxpool.Pool
}
