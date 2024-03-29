package main

import (
	"comies/api"
	"comies/api/handlers"
	"comies/api/middleware"
	"comies/app"
	"comies/config"
	"comies/data/conn"
	"comies/telemetry"
	"net"
	"os"
	"path"
	"strconv"
	"strings"
	"time"

	"github.com/bwmarrin/snowflake"
	"github.com/go-chi/chi/v5"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	cors "github.com/gorilla/handlers"
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

	logger.Info("Successfully created snowflake node")

	router := chi.NewRouter().With(middleware.Logging())
	handlers.Serve(router, handlers.Dependencies{
		App: app.NewApp(app.Deps{
			Snowflake: snflake,
		}),
		Pool: middleware.Pool(db),
		TX:   middleware.TX(db),
	})

	lis, err := net.Listen("tcp", cfg.Server.Address)
	if err != nil {
		logger.Fatal("Could not listen to port", zap.Error(err), zap.String("address", cfg.Server.Address))
	}

	api.Serve(lis, cors.CORS(
		cors.AllowCredentials(),
		cors.AllowedMethods([]string{"POST", "GET", "PUT", "DELETE", "OPTIONS"}),
		cors.AllowedHeaders([]string{"Content-Type"}),
		cors.AllowedOrigins(strings.Split(cfg.Server.CORSOrigins, ",")),
		cors.MaxAge(3600),
	)(router), time.Second*30, time.Second*30)

}
