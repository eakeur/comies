package order

import (
	"comies/app/core/entities/order"

	"github.com/jackc/pgx/v4/pgxpool"
)

var _ order.Actions = actions{}

type actions struct {
	db *pgxpool.Pool
}
