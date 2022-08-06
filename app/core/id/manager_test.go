package id

import (
	"comies/app/core/types"
	"testing"

	"github.com/bwmarrin/snowflake"
	"github.com/stretchr/testify/assert"
)

func Test_manager_Create(t *testing.T) {
	type args struct {
		id types.ID
	}
	type entity struct {
		id types.ID
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "should assign id",
			args: args{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			node, _ := snowflake.NewNode(1)
			manager{node: node}.Create(&tt.args.id)
			assert.Equal(t, tt.args.id != 0, true)
		})
	}
}
