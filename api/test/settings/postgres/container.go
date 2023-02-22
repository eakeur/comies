package postgres

import (
	"context"
	"fmt"

	"github.com/ory/dockertest"
	"github.com/ory/dockertest/docker"
)

const (
	containerName     = "postgres_testing_container"
	expires           = 600
	dockerPostgresURL = "postgres://postgres:postgres@localhost:%s/%s?sslmode=disable"
)

var (
	options = dockertest.RunOptions{
		Name:       containerName,
		Repository: "postgres",
		Tag:        "14-alpine",
		Env:        []string{"POSTGRES_USER=postgres", "POSTGRES_PASSWORD=postgres"},
	}
)

func Create(d *dockertest.Pool) (*dockertest.Resource, error) {
	container, _ := useContainer(d, containerName)
	if container != nil && container.Container.State.Running {
		return container, nil
	}

	if container != nil && !container.Container.State.Running {
		_ = d.RemoveContainerByName(containerName)
	}

	container, err := d.RunWithOptions(&options, func(c *docker.HostConfig) {
		c.AutoRemove = true
		c.RestartPolicy = docker.RestartPolicy{Name: "no"}
	})

	if err == nil {
		container.Expire(expires)
		return container, nil
	}

	err = d.Retry(func() error {
		container, err := useContainer(d, containerName)
		if container != nil {
			return nil
		}
		return err
	})

	if err != nil || container == nil {
		return nil, fmt.Errorf(`failed getting container resource: %w`, err)
	}

	return container, nil
}

func Startup(d *dockertest.Resource) (DatabaseContextBuilder, error) {
	port := d.GetPort("5432/tcp")

	conn, err := useConnection(url(port, "postgres"))
	if err != nil {
		return nil, fmt.Errorf("an error occurred when pinging database: %w", err)
	}

	database, ctx := createName(), context.Background()

	err = createDatabase(ctx, conn, database)
	if err != nil {
		return nil, err
	}

	conn.Close()

	conn, err = useConnection(url(port, database))
	if err != nil {
		return nil, fmt.Errorf("an error occurred when pinging database: %w", err)
	}

	err = mig(conn)
	if err != nil {
		return nil, fmt.Errorf("an error occurred migrating database: %w", err)
	}

	return builder(conn, port, database), createTemplate(ctx, conn, database)
}

func useContainer(d *dockertest.Pool, name string) (*dockertest.Resource, error) {
	container, err := d.Client.InspectContainer(name)
	if container == nil {
		return nil, err
	}

	return &dockertest.Resource{Container: container}, nil
}
