package ingredient

import (
	"github.com/jackc/pgx/v4/pgxpool"
	"gomies/app/core/entities/ingredient"
)

var _ ingredient.Actions = actions{}

type actions struct {
	db *pgxpool.Pool
}
