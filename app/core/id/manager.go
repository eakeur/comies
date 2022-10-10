package id

import (
	"comies/app/config"
	"strconv"

	"github.com/bwmarrin/snowflake"
)

type (
	ID int64
)

func (i ID) String() string {
	return strconv.FormatInt(int64(i), 10)
}

func (i ID) Empty() bool {
	return i == 0
}

var node *snowflake.Node

func NewNode(cfg config.IDGeneration) (err error) {
	number, err := strconv.Atoi(cfg.NodeNumber)
	if err != nil {
		return err
	}

	node, err = snowflake.NewNode(int64(number))
	return
}

func Create(id *ID) {
	gen := ID(node.Generate().Int64())
	*id = gen
}
