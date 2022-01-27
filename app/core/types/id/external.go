package id

import "github.com/google/uuid"

var Nil = External(uuid.Nil)

type External uuid.UUID

func ExternalFrom(input string) External {
	id, err := uuid.Parse(input)
	if err != nil {
		return Nil
	}
	return External(id)
}
