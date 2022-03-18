package types

import "github.com/google/uuid"

type (
	ID       int64
	External uuid.UUID
)

var Nil = External(uuid.Nil)

func ExternalFrom(input string) External {
	id, err := uuid.Parse(input)
	if err != nil {
		return Nil
	}
	return External(id)
}

func NewExternal() External {
	return External(uuid.New())
}

func (e External) Empty() bool {
	return e == External{} || e == Nil
}
