package listing

import (
	"github.com/stretchr/testify/assert"
	"gomies/app/core/entities/product"
	"gomies/app/core/types/entity"
	"testing"
)

func TestHasMore(t *testing.T) {
	t.Parallel()

	type test struct {
		name    string
		args    []entity.Entity
		appends []entity.Entity
		want    bool
	}

	tests := []test{
		{
			name: "should return has no more products",
			args: []entity.Entity{
				product.Product{
					Code: "", Name: "",
				}.Entity,
			},
			want: false,
		},
		{
			name: "should return has more products with slice",
			args: func() []entity.Entity {
				arr := []entity.Entity{
					product.Product{
						Code: "", Name: "",
					}.Entity,
					product.Product{
						Code: "", Name: "",
					}.Entity,
					product.Product{
						Code: "", Name: "",
					}.Entity,
				}
				return arr[0:2]
			}(),
			want: true,
		},
		{
			name: "should return has more products",
			args: make([]entity.Entity, 1, 10),
			appends: []entity.Entity{
				product.Product{
					Code: "", Name: "",
				}.Entity,
			},
			want: true,
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
