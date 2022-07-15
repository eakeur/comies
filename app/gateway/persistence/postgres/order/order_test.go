package order

import (
	"comies/app/gateway/persistence/postgres/tests"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	os.Exit(tests.SetupTest(m))
}
