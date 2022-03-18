package contacting

import "gomies/pkg/sdk/types"

type Address struct {
	types.Entity
	// TargetID is an identifier for the object this address references to
	TargetID   types.UID
	Code       string
	Street     string
	Number     string
	Complement string
	District   string
	City       string
	State      string
	Country    string
}
