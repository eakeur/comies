package test

import (
	"comies/app"
	"comies/core/types"
	"comies/io/data/postgres/tests"
	"comies/io/http"
	v1 "comies/io/http/handlers/v1"
	"comies/io/http/middleware"
	"comies/telemetry"
	"fmt"
	"os"
	"path"
	"testing"
	"time"

	"github.com/bwmarrin/snowflake"
	"github.com/go-chi/chi/v5"
	"go.uber.org/zap/zapcore"
)

func startup(t *testing.T) {
	t.Helper()

	logger := telemetry.NewLogger(os.Stdout, zapcore.DebugLevel)

	sqlLog, _ := os.Create(path.Join(os.TempDir(), fmt.Sprintf("%s_comies_sql_log_%d.log", t.Name(), time.Now().Nanosecond())))
	telemetry.Register(&telemetry.Telemetry{
		Logger: logger,
		SQL:    telemetry.NewLogger(sqlLog, zapcore.DebugLevel),
	})

	pool := tests.NewDBConn(t)

	snflake, err := snowflake.NewNode(21)
	if err != nil {
		t.Fatalf("Could not create snowflake node: %v", err)
	}

	var createID types.CreateID = func() types.ID {
		return types.ID(snflake.Generate())
	}

	router := chi.NewRouter().With(middleware.CORS(), middleware.Logging())
	v1.Serve(router, v1.Dependencies{
		App:  app.NewApp(createID),
		Pool: middleware.Pool(pool),
		TX:   middleware.TX(pool),
	})

	http.Serve(":5051", router, time.Second*30, time.Second*30)
}
