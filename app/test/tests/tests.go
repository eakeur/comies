package tests

import (
	"comies/app"
	"comies/app/gateway/api"
	servers "comies/app/gateway/api/tests"
	database "comies/app/gateway/persistence/postgres/tests"
	"context"
	"github.com/bwmarrin/snowflake"
	"google.golang.org/grpc"
	"log"
	"testing"
)

func SetupTest(m *testing.M) int {
	return database.SetupTest(m)
}

func NewTestApp(t *testing.T) *grpc.ClientConn {
	ctx := context.Background()
	db := database.NewTestDatabase(t, ctx)

	flake, err := snowflake.NewNode(1)
	if err != nil {
		log.Fatalf("Could not create snowflake node: %v", err)
	}

	application := app.NewApplication(app.Gateways{
		Database:      db.Pool,
		SnowflakeNode: flake,
	})

	srv := api.NewAPI(application)

	t.Cleanup(func() {
		srv.GracefulStop()
		db.Drop()
	})

	return servers.NewTestServer(t, srv)
}
