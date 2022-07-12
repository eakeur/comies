package member

import (
	"github.com/jackc/pgx/v4/pgxpool"
	"gomies/app/core/entities/member"
)

var _ member.Actions = actions{}

type actions struct {
	db *pgxpool.Pool
}
