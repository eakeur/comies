package listing

import (
	"github.com/stretchr/testify/assert"
	"gomies/pkg/sdk/types"
	"testing"
)

func TestHasMore(t *testing.T) {
	t.Parallel()

	type test struct {
		name    string
		args    []types.Entity
		appends []types.Entity
		want    bool
	}

	tests := []test{
		{
			name: "should return has no more products",
			args: []types.Entity{
				{},
			},
			want: false,
		},
		{
			name: "should return has more products with slice",
			args: func() []types.Entity {
				arr := []types.Entity{{}, {}, {}}
				return arr[0:2]
			}(),
			want: true,
		},
		{
			name:    "should return has more products",
			args:    make([]types.Entity, 1, 10),
			appends: []types.Entity{{}},
			want:    true,
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			if tc.appends != nil {
				for _, a := range tc.appends {
					tc.args = append(tc.args, a)
				}
			}

			assert.Equal(t, tc.want, HasMore(len(tc.args), cap(tc.args)))
		})
	}
}
