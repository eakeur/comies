package movement

import (
	"github.com/jackc/pgx/v4/pgxpool"
	"gomies/app/core/entities/movement"
)

var _ movement.Actions = actions{}

type actions struct {
	db *pgxpool.Pool
}
