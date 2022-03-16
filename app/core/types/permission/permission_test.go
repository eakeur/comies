package permission

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPermissions_Allow(t *testing.T) {
	t.Parallel()

	type test struct {
		name      string
		perm      Permissions
		operation string
		wantErr   error
	}

	cases := []test{
		{
			name:      "should return nil for authorized",
			operation: "Workflows.Product.Delete",
			perm:      "Workflows.Product.Add;Workflows.Product.Delete",
		},
		{
			name:      "should return err for unauthorized",
			operation: "Workflows.Product.Add",
			wantErr:   ErrNotAllowed,
			perm:      "Workflows.Product.Delete",
		},
		{
			name:      "should return nil for wildcard",
			operation: "Workflows.Product.Add",
			perm:      "*",
		},
	}

	for _, tc := range cases {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			assert.ErrorIs(t, tc.perm.CheckAllowance(tc.operation), tc.wantErr)
		})
	}
}
