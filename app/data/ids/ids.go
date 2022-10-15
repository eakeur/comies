package ids

import (
	"comies/app/config"
	"comies/app/core/id"
	"github.com/bwmarrin/snowflake"
	"strconv"
)

var node *snowflake.Node

func NewNode(cfg config.IDGeneration) (err error) {
	number, err := strconv.Atoi(cfg.NodeNumber)
	if err != nil {
		return err
	}

	node, err = snowflake.NewNode(int64(number))
	return
}

func Create() id.ID {
	return id.ID(node.Generate().Int64())
}