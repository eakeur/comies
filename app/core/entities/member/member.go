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
	}
)

func (m Member) Validate() error {
	// TODO implement validation
	return nil
}
