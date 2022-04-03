package contacting

import (
	"context"
	"github.com/stretchr/testify/assert"
	"gomies/app/core/entities/contacting"
	"gomies/pkg/sdk/types"
	"testing"
)

func TestWorkflow_SaveProduct(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	type (
		args struct {
			contact contacting.Contact
		}

		opts struct {
			contacts *contacting.ActionsMock
		}

		test struct {
			name    string
			args    args
			opts    opts
			want    contacting.Contact
			wantErr error
		}
	)

	cases := []test{
		{
			name: "should return created contact",
			args: args{
				contact: contacting.Contact{
					TargetID: idExample1,
					Phones: []contacting.Phone{
						{}, {}, {}, {}, {}, {}, {}, {}, {}, {},
					},
					Addresses: []contacting.Address{
						{}, {}, {}, {}, {}, {}, {}, {}, {}, {},
					},
				},
			},
			want: contacting.Contact{
				TargetID: idExample1,
				Phones: []contacting.Phone{
					{}, {}, {}, {}, {}, {}, {}, {}, {}, {},
				},
				Addresses: []contacting.Address{
					{}, {}, {}, {}, {}, {}, {}, {}, {}, {},
				},
			},
			opts: opts{
				contacts: &contacting.ActionsMock{
					SaveAddressesFunc: func(ctx context.Context, target types.UID, addresses ...contacting.Address) ([]contacting.Address, error) {
						return addresses, nil
					},
					SavePhonesFunc: func(ctx context.Context, target types.UID, phones ...contacting.Phone) ([]contacting.Phone, error) {
						return phones, nil
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
			ingredient, err := wf.SaveContact(ctx, c.args.contact)

			assert.ErrorIs(t, err, c.wantErr)
			assert.Equal(t, c.want, ingredient)

		})
	}

}
