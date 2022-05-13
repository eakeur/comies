package crew

import (
	"gomies/app/sdk/types"
	"time"
)

type (
	Member struct {
		types.Entity
		Name           string
		FullName       string
		Nickname       string
		Reference      string
		PictureURL     string
		PasswordChange time.Time
		LastSeen       time.Time
		Password       types.Password
		Permissions    types.Permissions
		types.Store
	}

	MemberType int
)

func (m Member) Validate() error {
	// TODO implement validation
	return nil
}
