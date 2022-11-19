package ingredient

import (
	"comies/app/core/menu/ingredient"

	"github.com/jackc/pgx/v4/pgxpool"
)

var _ ingredient.Actions = actions{}

type actions struct {
	db *pgxpool.Pool
}

func NewActions(db *pgxpool.Pool) ingredient.Actions {
	return &actions{db: db}
}
