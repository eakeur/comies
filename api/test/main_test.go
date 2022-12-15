package test

import (
	"comies/io/data/postgres/tests"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	os.Exit(tests.SetupTest(m))
}
