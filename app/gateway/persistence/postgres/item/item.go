package item

import (
	"github.com/jackc/pgx/v4/pgxpool"
	"gomies/app/core/entities/item"
)

var _ item.Actions = actions{}

type actions struct {
	db *pgxpool.Pool
}
