package main

import (
	"comies/config"
	"comies/io/data/conn"
	"comies/telemetry"
	"errors"
	"os"
	"strconv"

	migrate "github.com/golang-migrate/migrate/v4"
	cli "github.com/urfave/cli/v2"
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

	telemetry.Register(&telemetry.Telemetry{
		Logger: logger,
		SQL:    telemetry.NewLogger(os.Stdout, zapcore.PanicLevel),
	})

	db, err := conn.Connect(cfg.Database)
	if err != nil {
		logger.Fatal("Could not connect postgres database", zap.Error(err))
	}

	mig, err := conn.MigrationHandler(db)
	if err != nil {
		logger.Fatal("Could not create migration handler", zap.Error(err))
	}

	app := &cli.App{
		Commands: []*cli.Command{
			{
				Name:    "migrate",
				Aliases: []string{"m"},
				Usage:   "migrate commands",
				Subcommands: []*cli.Command{
					{
						Name:  "up",
						Usage: "migrate up",
						Action: func(c *cli.Context) error {
							return mig.Up()
						},
					},
					{
						Name:  "down",
						Usage: "migrate down",
						Action: func(c *cli.Context) error {
							return mig.Down()
						},
					},
					{
						Name:  "force",
						Usage: "migrate force",
						Action: func(c *cli.Context) error {
							version, _ := strconv.Atoi(c.Args().First())
							return mig.Force(version)
						},
					},
				},
			},
		},
	}

	app.Name = "Comies CLI"
	app.Usage = "migration tooling"

	err = app.Run(os.Args)
	if err != nil {
		switch {
		case errors.Is(err, migrate.ErrNoChange):
		case errors.Is(err, migrate.ErrNilVersion):
			logger.Info(err.Error())
		default:
			logger.Fatal(err.Error())
		}
	}
}
