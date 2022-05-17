package phone

import "gomies/app/sdk/types"

type Phone struct {
	ID          types.ID
	History     types.History
	TargetID    types.ID
	CountryCode string
	AreaCode    string
	Number      string
	Active      bool
}
