package test

import (
	"comies/api"
	v1 "comies/api/handlers/v1"
	"comies/api/middleware"
	"comies/app"
	"comies/telemetry"
	"comies/test/settings/postgres"
	"fmt"
	"net"
	"os"
	"testing"
	"time"

	"github.com/bwmarrin/snowflake"
	"github.com/go-chi/chi/v5"
	"go.uber.org/zap/zapcore"
)

var db postgres.DatabaseContextBuilder

func createAPI(t *testing.T) string {
	t.Helper()

	logger := telemetry.NewLogger(os.Stdout, zapcore.DebugLevel)

	telemetry.Register(&telemetry.Telemetry{
		Logger: logger,
		SQL:    telemetry.NewLogger(os.Stdout, zapcore.WarnLevel),
	})

	pool := db.Connection(t)

	snflake, err := snowflake.NewNode(21)
	if err != nil {
		t.Fatalf("Could not create snowflake node: %v", err)
	}

	router := chi.NewRouter().With(middleware.CORS(), middleware.Logging())
	v1.Serve(router, v1.Dependencies{
		App: app.NewApp(app.Deps{
			Snowflake: snflake,
		}),
		Pool: middleware.Pool(pool),
		TX:   middleware.TX(pool),
	})

	lis, err := net.Listen("tcp", "localhost:")
	if err != nil {
		t.Fatal(err)
	}

	go api.Serve(lis, router, time.Second*30, time.Second*30)

	return fmt.Sprintf("http://%s", lis.Addr().String())
}
