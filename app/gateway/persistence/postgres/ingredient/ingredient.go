package ingredient

import (
	"comies/app/core/entities/ingredient"

	"github.com/jackc/pgx/v4/pgxpool"
)

var _ ingredient.Actions = actions{}

type actions struct {
	db *pgxpool.Pool
}

func NewRepository(db *pgxpool.Pool) ingredient.Actions {
	return actions{db: db}
}
