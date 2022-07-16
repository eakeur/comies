package menu

import (
	"comies/app/gateway/api/gen/menu"
	"comies/app/test/tests"
	"context"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	os.Exit(tests.SetupTest(m))
}

func NewClient(t *testing.T) (context.Context, menu.MenuClient) {
	return context.Background(), menu.NewMenuClient(tests.NewTestApp(t))
}
