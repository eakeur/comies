package product

import (
	"comies/app/core/entities/product"

	"github.com/jackc/pgx/v4/pgxpool"
)

var _ product.Actions = actions{}

type actions struct {
	db *pgxpool.Pool
}

func NewRepository(db *pgxpool.Pool) product.Actions {
	return actions{db: db}
}
