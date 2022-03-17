package crew

import (
	"gomies/pkg/sdk/types"
	"time"
)

type (
	Member struct {
		types.Entity
		FirstName       string
		LastName        string
		Nickname        string
		Password        string
		Reference       string
		PictureURL      string
		PasswordChange  time.Time
		LastSeen        time.Time
		Permissions     types.Permissions
		StoreID         types.ID
		StoreExternalID types.External
	}

	MemberType int
)

func (m Member) Validate() error {
	// TODO implement validation
	return nil
}
