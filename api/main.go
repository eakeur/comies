package main

import (
	"comies/config"
	"comies/core/types"
	"comies/io/data/postgres/conn"
	"comies/io/data/postgres/ingredient"
	"comies/io/data/postgres/item"
	"comies/io/data/postgres/movement"
	"comies/io/data/postgres/order"
	"comies/io/data/postgres/price"
	"comies/io/data/postgres/product"
	"comies/io/data/postgres/status"
	"comies/io/http"
	v1 "comies/io/http/handlers/v1"
	"comies/io/http/middleware"
	"comies/jobs/menu"
	"comies/jobs/ordering"
	"comies/telemetry"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/bwmarrin/snowflake"
	"github.com/go-chi/chi/v5"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func main() {
	cfg := config.Load()

	logger := telemetry.NewLogger(os.Stdout, map[string]zapcore.Level{
		"production":  zapcore.InfoLevel,
		"development": zapcore.DebugLevel,
	}["production"])

	telemetry.Register(&telemetry.Telemetry{Logger: logger})

	logger.Info("Successfully created logger instance")

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
		log.Fatalf("Could not parse id generation node number %v: %v", cfg.IDGeneration.NodeNumber, err)
	}

	snflake, err := snowflake.NewNode(int64(nodeNumber))
	if err != nil {
		log.Fatalf("Could not create snowflake node: %v", err)
	}

	createID := func() types.ID {
		return types.ID(snflake.Generate())
	}

	log.Println("Successfully created snowflake node")

	menu := menu.NewJobs(product.NewActions(), ingredient.NewActions(), movement.NewActions(), price.NewActions(), createID)
	ordering := ordering.NewJobs(order.NewActions(), item.NewActions(), status.NewActions(), createID, menu.GetProductLatestPriceByID, nil)

	router := chi.NewRouter().With(middleware.CORS(), middleware.Logging())
	v1.Serve(router, v1.Dependencies{
		Menu:     menu,
		Ordering: ordering,
		Pool:     middleware.Pool(db),
		TX:       middleware.TX(db),
	})

	http.Serve(cfg.Server.ListenPort, router, time.Second*30, time.Second*30)

}
