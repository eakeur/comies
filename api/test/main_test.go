package test

import (
	"comies/test/settings"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	docker, err := settings.NewPool()
	if err != nil {
		panic(err.Error())
	}

	d, purge, err := settings.Postgres(docker)
	if err != nil {
		panic(err.Error())
	}

	db = d
	defer purge()

	os.Exit(m.Run())
}
