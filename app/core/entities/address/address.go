package address

import "gomies/app/sdk/types"

type Address struct {
	ID         types.ID
	History    types.History
	TargetID   types.ID
	Code       string
	Street     string
	Number     string
	Complement string
	District   string
	City       string
	State      string
	Country    string
	Active     bool
}
