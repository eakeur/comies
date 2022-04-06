package crew

import (
	"context"
	"gomies/app/core/entities/iam/crew"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWorkflow_UpdateMember(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	type (
		args struct {
			member crew.Member
		}

		opts struct {
			crew *crew.ActionsMock
		}

		test struct {
			name    string
			args    args
			opts    opts
			wantErr error
		}
	)

	cases := []test{
		{
			name: "should return nil for successful update",
			args: args{
				member: crew.Member{},
			},
			opts: opts{
				crew: &crew.ActionsMock{
					UpdateMemberFunc: func(ctx context.Context, op crew.Member) error {
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

			wf := workflow{
				crew: c.opts.crew,
			}
			err := wf.UpdateMember(ctx, c.args.member)
			assert.ErrorIs(t, err, c.wantErr)
		})
	}

}
