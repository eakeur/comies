package product

import (
	"github.com/jackc/pgx/v4/pgxpool"
	"gomies/app/core/entities/product"
)

var _ product.Actions = actions{}

type actions struct {
	db *pgxpool.Pool
}
