package session

import (
	"gomies/pkg/sdk/types"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPermissions_Allow(t *testing.T) {
	t.Parallel()

	type test struct {
		name      string
		perm      types.Permissions
		operation string
		want      bool
	}

	cases := []test{
		{
			name:      "should return true for authorized",
			operation: "Workflows.Product.Delete",
			perm:      "Workflows.Product.Add;Workflows.Product.Delete",
			want:      true,
		},
		{
			name:      "should return err for unauthorized",
			operation: "Workflows.Product.Add",
			perm:      "Workflows.Product.Delete",
		},
		{
			name:      "should return true for wildcard",
			operation: "Workflows.Product.Add",
			perm:      "*",
			want:      true,
		},
	}

	for _, tc := range cases {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			assert.Equal(t, Session{Permissions: tc.perm}.isAllowed(tc.operation), tc.want)
		})
	}
}
