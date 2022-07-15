package movement

import (
	"comies/app/core/entities/movement"

	"github.com/jackc/pgx/v4/pgxpool"
)

var _ movement.Actions = actions{}

type actions struct {
	db *pgxpool.Pool
}

func NewRepository(db *pgxpool.Pool) movement.Actions {
	return actions{db: db}
}
