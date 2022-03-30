package contacting

import (
	"context"
	"github.com/stretchr/testify/assert"
	"gomies/pkg/contacting/core/entities/contacting"
	"gomies/pkg/sdk/types"
	"testing"
)

func TestWorkflow_ListPhones(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	type (
		args struct {
			targetID types.UID
		}

		opts struct {
			contacts *contacting.ActionsMock
		}

		test struct {
			name    string
			args    args
			opts    opts
			want    []contacting.Phone
			wantErr error
		}
	)

	cases := []test{
		{
			name: "should return addresses",
			args: args{
				targetID: idExample1,
			},
			want: []contacting.Phone{},
			opts: opts{
				contacts: &contacting.ActionsMock{
					ListPhonesFunc: func(ctx context.Context, target types.UID) ([]contacting.Phone, error) {
						return []contacting.Phone{}, nil
					},
				},
			},
		},
	}

	for _, c := range cases {
		c := c

		t.Run(c.name, func(t *testing.T) {
			t.Parallel()

			wf := NewWorkflow(c.opts.contacts)
			got, err := wf.ListPhones(ctx, c.args.targetID)

			assert.ErrorIs(t, err, c.wantErr)
			assert.Equal(t, c.want, got)

		})
	}
}
