package movement

import (
	"comies/core/menu/movement"
	"comies/test/settings/postgres"
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBalance(t *testing.T) {
	t.Parallel()

	type args struct {
		filter movement.Filter
	}

	for _, tt := range []struct {
		name         string
		args         args
		checkBalance assert.ValueAssertionFunc
		checkErr     assert.ErrorAssertionFunc
		before       postgres.Callback
	}{
		{
			name:         "should return 0 as balance",
			checkBalance: assert.Zero,
			checkErr:     assert.NoError,
			before: func(ctx context.Context, t *testing.T) {

			},
			args: args{
				filter: movement.Filter{
					ProductID: 838737463,
				},
			},
		},
	} {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			ctx := db.Pool(t, tt.before)
			bal, err := actions{}.Balance(ctx, tt.args.filter)

			tt.checkErr(t, err)
			tt.checkBalance(t, bal)
		})
	}
}
