package types

import "github.com/google/uuid"

type (
	ID  int64
	UID string
)

func NewUID() UID {
	return UID(uuid.NewString())
}

func (e UID) Empty() bool {
	return e == ""
}
