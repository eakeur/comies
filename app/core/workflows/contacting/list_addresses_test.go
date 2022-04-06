package contacting

import (
	"context"
	"gomies/app/core/entities/contacting"
	"gomies/pkg/sdk/types"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWorkflow_ListAddress(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	fakeID := types.UID("1bdcafba-9deb-48b4-8a0e-ecea4c99b0e3")

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
			want    []contacting.Address
			wantErr error
		}
	)

	cases := []test{
		{
			name: "should return addresses",
			args: args{
				targetID: fakeID,
			},
			want: []contacting.Address{},
			opts: opts{
				contacts: &contacting.ActionsMock{
					ListAddressesFunc: func(ctx context.Context, target types.UID) ([]contacting.Address, error) {
						return []contacting.Address{}, nil
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
			got, err := wf.ListAddresses(ctx, c.args.targetID)

			assert.ErrorIs(t, err, c.wantErr)
			assert.Equal(t, c.want, got)

		})
	}
}
