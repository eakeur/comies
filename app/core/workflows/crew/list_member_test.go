package crew

import (
	"context"
	"gomies/app/core/entities/iam/crew"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWorkflow_List(t *testing.T) {
	t.Parallel()

	type (
		args struct {
			filter crew.Filter
		}

		fields struct {
			crew *crew.ActionsMock
		}

		test struct {
			name    string
			args    args
			fields  fields
			want    []crew.Member
			wantErr error
		}
	)

	cases := []test{
		{
			name: "should return member array",
			args: args{
				filter: crew.Filter{},
			},
			want: []crew.Member{},
			fields: fields{
				crew: &crew.ActionsMock{
					ListMembersFunc: func(ctx context.Context, operatorFilter crew.Filter) ([]crew.Member, int, error) {
						return []crew.Member{}, 0, nil
					},
				},
			},
		},
	}

	for _, c := range cases {
		c := c

		t.Run(c.name, func(t *testing.T) {
			t.Parallel()

			list, _, err := workflow{crew: c.fields.crew}.ListMembers(context.Background(), c.args.filter)
			assert.ErrorIs(t, err, c.wantErr)
			assert.Equal(t, c.want, list)

		})
	}
}
