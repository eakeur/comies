package movement

import (
	"comies/core/menu/movement"
	"comies/data/conn"
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
			args: args{
				filter: movement.Filter{
					ProductID: 838737463,
				},
			},
		},
		{
			name:         "should return 30 as balance",
			checkBalance: assert.Zero,
			checkErr:     assert.NoError,
			args: args{
				filter: movement.Filter{
					ProductID: 838737463,
				},
			},
			before: func(ctx context.Context, t *testing.T) {
				const script = `
					insert into products (
						id,
						code,
						name,
						type,
						cost_price,
						sale_unit,
						minimum_sale,
						minimum_quantity,
						maximum_quantity,
						location
					) values (
						1, 'cod', 'name', 10, 2, 'un', 1, 1, 10, ''
					);

					insert into movements (
						id,
						product_id,
						type,
						date,
						quantity,
						agent_id
					) values (
						1, 1, 2, now(), 30, 1
					);
				`

				_, err := conn.ExecFromContext(ctx, script)
				if err != nil {
					t.Fatal(err)
				}
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
