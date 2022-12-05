package tests

import (
	"comies/config"
	"comies/io/data/postgres/conn"
	"context"
	"crypto/rand"
	"fmt"
	"log"
	"math"
	"math/big"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/ory/dockertest/docker"

	"github.com/ory/dockertest"
	"github.com/pkg/errors"
)

const (
	ContainerName    = "postgres_testing_container"
	ContainerExpires = 600
)

func ConnectToDockerPostgres() (*pgxpool.Pool, error) {
	pool, err := pool()
	resource, err := resource(pool)
	if err != nil {
		return nil, errors.Wrap(err, "the docker container resource could not be created or fetched")
	}

	dbPortBinding := resource.GetPort("5432/tcp")
	if err != nil {
		return nil, errors.Wrap(err, "could not set an expiration time")
	}

	db, err := connect(pool, dbPortBinding)
	if err != nil {
		defer pool.Purge(resource)
		return nil, errors.Wrap(err, "an error occurred when pinging database")
	}

	n, _ := rand.Int(rand.Reader, big.NewInt(math.MaxInt32))
	databaseName := fmt.Sprintf("database_%d", n)

	err = up(db, databaseName)
	if err != nil {
		defer pool.Purge(resource)
		return nil, err
	}

	Connection, err = conn.Connect(config.Database{
		User:     "postgres",
		Password: "postgres",
		Host:     "localhost:" + dbPortBinding,
		Name:     databaseName,
		SSLMode:  "disable",
	})
	if err != nil {
		defer pool.Purge(resource)
		return nil, errors.Wrap(err, "an error occurred when creating test database schema")
	}

	err = conn.Migrate(Connection)
	if err != nil {
		defer pool.Purge(resource)
		return nil, errors.Wrap(err, "an error occurred when migrating test database schema")
	}

	_, _ = Connection.Exec(
		context.Background(), fmt.Sprintf(
			"create database %[1]s_template template %[1]s;",
			databaseName,
		),
	)

	return Connection, nil

}

func pool() (*dockertest.Pool, error) {
	pool, err := dockertest.NewPool("")
	if err != nil {
		return nil, errors.Wrap(err, "the docker pool connection could not be established")
	}

	if err := pool.Client.Ping(); err != nil {
		return nil, errors.Wrap(err, "could not contact docker pool")
	}

	return pool, nil
}

func resource(dockerPool *dockertest.Pool) (*dockertest.Resource, error) {
	container, _ := dockerPool.Client.InspectContainer(ContainerName)
	if container != nil && container.State.Running {
		resource := &dockertest.Resource{Container: container}
		return resource, nil
	}
	if container != nil && !container.State.Running {
		_ = dockerPool.RemoveContainerByName(ContainerName)
	}

	resource, err := dockerPool.RunWithOptions(&dockertest.RunOptions{
		Name:       ContainerName,
		Repository: "postgres",
		Tag:        "14-alpine",
		Env:        []string{"POSTGRES_USER=postgres", "POSTGRES_PASSWORD=postgres"},
	}, func(c *docker.HostConfig) {
		c.AutoRemove = true
		c.RestartPolicy = docker.RestartPolicy{Name: "no"}
	})
	if err != nil {
		log.Println("concurrent processes from different packages might be starting docker, try to connect to them")
		var container *docker.Container
		if err = dockerPool.Retry(func() error {
			if container, err = dockerPool.Client.InspectContainer(ContainerName); container != nil {
				return nil
			}
			return err
		}); err != nil || container == nil {
			return nil, fmt.Errorf(`failed getting container resource: %w`, err)
		}

		log.Println("Got connection to docker from concurrent process")
		return &dockertest.Resource{Container: container}, nil
	}

	resource.Expire(ContainerExpires)
	return resource, nil
}

func connect(pool *dockertest.Pool, at string) (*pgxpool.Pool, error) {
	var pg *pgxpool.Pool

	if err := pool.Retry(func() error {
		var err error
		pg, err = conn.Connect(config.Database{
			User:     "postgres",
			Password: "postgres",
			Host:     "localhost:" + at,
			Name:     "postgres",
			SSLMode:  "disable",
		})
		if err != nil {
			return err
		}

		return pg.Ping(context.Background())
	}); err != nil {
		return nil, errors.Wrap(err, "an error occurred when connecting to the database server")
	}

	return pg, nil
}

func up(conn *pgxpool.Pool, name string) error {
	_ = down(conn, name)

	if _, err := conn.Exec(context.Background(), fmt.Sprintf(`create database %s;`, name)); err != nil {
		return fmt.Errorf("failed creating test database %s: %w", name, err)
	}
	return nil
}

func down(conn *pgxpool.Pool, name string) error {
	if _, err := conn.Exec(context.Background(), fmt.Sprintf(`drop database if exists %s;`, name)); err != nil {
		return fmt.Errorf("failed dropping test database %s: %w", name, err)
	}

	return nil
}
