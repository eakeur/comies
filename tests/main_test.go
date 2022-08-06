package tests

import (
	"comies/app"
	"comies/app/gateway/api"
	database "comies/app/gateway/persistence/postgres/tests"
	"context"
	"fmt"
	"github.com/bwmarrin/snowflake"
	"go.uber.org/zap"
	"log"
	"net"
	"net/http"
	"net/url"
	"testing"
	"time"
)

func TestMain(m *testing.M) {
	database.SetupTest(m)
}

func NewTestApp(t *testing.T, defaultRoute string) Client {
	ctx := context.Background()
	db := database.NewTestDatabase(t, ctx)

	flake, err := snowflake.NewNode(1)
	if err != nil {
		log.Fatalf("Could not create snowflake node: %v", err)
	}

	logger, err := zap.NewDevelopment()
	if err != nil {
		log.Fatalf("Could not create logger instance: %v", err)
	}

	application := app.NewApplication(app.Gateways{
		Logger:        logger.Sugar(),
		Database:      db.Pool,
		SnowflakeNode: flake,
	})

	lis, err := net.Listen("tcp", "localhost:")
	if err != nil {
		log.Fatalf("Could not open test server: %v", err)
	}

	addr := lis.Addr()

	t.Cleanup(func() {
		db.Drop()
		_ = lis.Close()
	})

	go http.Serve(lis, api.NewAPI(application))

	origin, err := url.Parse(fmt.Sprintf("http://%s%s", addr, defaultRoute))
	if err != nil {
		log.Fatalf("Failed mounting URL: %v", err)
	}

	return client{
		client: &http.Client{
			Timeout: time.Second * 120,
		},
		origin: origin,
	}
}
