package id

import (
	"comies/app/core/types"
	"github.com/bwmarrin/snowflake"
)

type (
	Manager interface {
		Create(id *types.ID)
	}

	manager struct {
		node *snowflake.Node
	}
)

func NewManager(node *snowflake.Node) Manager {
	return manager{
		node: node,
	}
}

func (m manager) Create(id *types.ID) {
	gen := types.ID(m.node.Generate().Int64())
	*id = gen
}
