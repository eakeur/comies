package tests

import (
	"github.com/jackc/pgx/v4/pgxpool"
	"os"
	"testing"
)

var conn *pgxpool.Pool

func SetupTest(m *testing.M) int {

	c, err := ConnectToDockerPostgres()
	if err != nil {
		os.Exit(1)
	}

	conn = c
	return m.Run()
}
