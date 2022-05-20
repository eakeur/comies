package ingredient

import (
	"gomies/app/gateway/persistence/postgres/tests"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	os.Exit(tests.SetupTest(m))
}
