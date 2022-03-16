package contacting

import (
	"gomies/app/core/types/entity"
	"gomies/app/core/types/id"
)

type Phone struct {
	entity.Entity

	// TargetID is an identifier for the object this phone references to
	TargetID id.External

	CountryCode string
	AreaCode    string
	Number      string
}
