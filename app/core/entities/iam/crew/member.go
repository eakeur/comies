package crew

import (
	"gomies/pkg/sdk/types"
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
		StoreID        types.ID
		types.Store
	}

	MemberType int
)

func (m Member) Validate() error {
	// TODO implement validation
	return nil
}
