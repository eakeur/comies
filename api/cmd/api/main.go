package main

import (
	"comies/app"
	"comies/config"
	"comies/core/types"
	"comies/io/data/postgres/conn"
	"comies/io/http"
	v1 "comies/io/http/handlers/v1"
	"comies/io/http/middleware"
	"comies/telemetry"
	"os"
	"path"
	"strconv"
	"time"

	"github.com/bwmarrin/snowflake"
	"github.com/go-chi/chi/v5"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func main() {
	cfg := config.Load()

	level := map[string]zapcore.Level{
		"production":  zapcore.InfoLevel,
		"development": zapcore.DebugLevel,
	}[cfg.Logger.Environment]

	logger := telemetry.NewLogger(os.Stdout, level)

	sqlLogFilePath := path.Join(os.TempDir(), "comies_sql.log")

	sqlLog, _ := os.Create(sqlLogFilePath)
	telemetry.Register(&telemetry.Telemetry{
		Logger: logger,
		SQL:    telemetry.NewLogger(sqlLog, level),
	})

	logger.Info("Successfully created logger instance", zap.String("sql", sqlLogFilePath))

	db, err := conn.Connect(cfg.Database)
	if err != nil {
		logger.Fatal("Could not connect postgres database", zap.Error(err))
	}

	err = conn.Migrate(db)
	if err != nil {
		logger.Fatal("Could not migrate postgres database", zap.Error(err))
	}

	logger.Info("Successfully connected and migrated database", zap.String("database", db.Config().ConnConfig.Database))

	nodeNumber, err := strconv.Atoi(cfg.IDGeneration.NodeNumber)
	if err != nil {
		logger.Fatal("Could not parse id generation node number", zap.String("node", cfg.IDGeneration.NodeNumber), zap.Error(err))
	}

	snflake, err := snowflake.NewNode(int64(nodeNumber))
	if err != nil {
		logger.Fatal("Could not create snowflake node", zap.Error(err))
	}

	var createID types.CreateID = func() types.ID {
		return types.ID(snflake.Generate())
	}

	logger.Info("Successfully created snowflake node")

	router := chi.NewRouter().With(middleware.CORS(), middleware.Logging())
	v1.Serve(router, v1.Dependencies{
		App:  app.NewApp(createID),
		Pool: middleware.Pool(db),
		TX:   middleware.TX(db),
	})

	http.Serve(cfg.Server.ListenPort, router, time.Second*30, time.Second*30)

}
