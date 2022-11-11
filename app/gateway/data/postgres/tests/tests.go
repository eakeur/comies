package tests

import (
	"os"
	"testing"

	"github.com/jackc/pgx/v4/pgxpool"
)

var Connection *pgxpool.Pool

func SetupTest(m *testing.M) int {

	c, err := ConnectToDockerPostgres()
	if err != nil {
		os.Exit(1)
	}

	Connection = c
	return m.Run()
}
