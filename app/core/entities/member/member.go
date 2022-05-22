package member

import (
	"gomies/app/sdk/types"
)

type (
	Member struct {
		ID       types.ID
		Active   bool
		Name     string
		Nickname string
		Password types.Password
		History  types.History
	}

	MemberType int
)

func (m Member) Validate() error {
	// TODO implement validation
	return nil
}
