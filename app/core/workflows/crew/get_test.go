package crew

import (
	"context"
	"github.com/stretchr/testify/assert"
	"gomies/app/core/entities/iam/crew"
	"gomies/pkg/sdk/fault"
	"gomies/pkg/sdk/types"
	"testing"
)

func TestWorkflow_Get(t *testing.T) {
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
			want    crew.Member
			wantErr error
		}
	)

	cases := []test{
		{
			name: "should return member found",
			args: args{key: crew.Key{ID: types.NewUID()}},
			fields: fields{
				crew: &crew.ActionsMock{
					GetFunc: func(ctx context.Context, key crew.Key) (crew.Member, error) {
						return crew.Member{}, nil
					},
				},
			},
		},
		{
			name:    "should return error not found",
			args:    args{key: crew.Key{ID: types.NewUID()}},
			wantErr: fault.ErrNotFound,
			fields: fields{
				crew: &crew.ActionsMock{
					GetFunc: func(ctx context.Context, key crew.Key) (crew.Member, error) {
						return crew.Member{}, fault.ErrNotFound
					},
				},
			},
		},
	}

	for _, c := range cases {
		c := c

		t.Run(c.name, func(t *testing.T) {
			t.Parallel()

			mem, err := workflow{crew: c.fields.crew}.Get(context.Background(), c.args.key)
			assert.ErrorIs(t, err, c.wantErr)
			assert.Equal(t, c.want, mem)

		})
	}
}
