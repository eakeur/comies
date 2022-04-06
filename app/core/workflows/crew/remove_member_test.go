package crew

import (
	"context"
	"gomies/app/core/entities/iam/crew"
	"gomies/pkg/sdk/types"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWorkflow_Remove(t *testing.T) {
	t.Parallel()

	type (
		args struct {
			key crew.Key
		}

		fields struct {
			crew *crew.ActionsMock
		}

		test struct {
			name    string
			args    args
			fields  fields
			wantErr error
		}
	)

	cases := []test{
		{
			name: "should return member found",
			args: args{key: crew.Key{ID: types.NewUID()}},
			fields: fields{
				crew: &crew.ActionsMock{
					RemoveMemberFunc: func(ctx context.Context, key crew.Key) error {
						return nil
					},
				},
			},
		},
	}

	for _, c := range cases {
		c := c

		t.Run(c.name, func(t *testing.T) {
			t.Parallel()

			err := workflow{crew: c.fields.crew}.RemoveMember(context.Background(), c.args.key)
			assert.ErrorIs(t, err, c.wantErr)

		})
	}
}
