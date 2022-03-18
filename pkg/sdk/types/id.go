package types

import "github.com/google/uuid"

type (
	ID  int64
	UID uuid.UUID
)

var Nil = UID(uuid.Nil)

func UIDFrom(input string) UID {
	id, err := uuid.Parse(input)
	if err != nil {
		return Nil
	}
	return UID(id)
}

func NewUID() UID {
	return UID(uuid.New())
}

func (e UID) Empty() bool {
	return e == UID{} || e == Nil
}
