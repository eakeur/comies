package contacting

import (
	"gomies/app/core/types/entity"
	"gomies/app/core/types/id"
)

type Address struct {
	entity.Entity
	// TargetID is an identifier for the object this address references to
	TargetID   id.External
	Code       string
	Street     string
	Number     string
	Complement string
	District   string
	City       string
	State      string
	Country    string
}
