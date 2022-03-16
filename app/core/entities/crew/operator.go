package crew

import (
	"gomies/app/core/types/entity"
	"gomies/app/core/types/permission"
	"time"
)

type Operator struct {
	entity.Entity
	FirstName      string
	LastName       string
	Nick           string
	Password       string
	PasswordChange time.Time
	LastSeen       time.Time
	Permissions    permission.Permissions
}
