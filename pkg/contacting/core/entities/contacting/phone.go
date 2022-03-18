package contacting

import "gomies/pkg/sdk/types"

type Phone struct {
	types.Entity

	// TargetID is an identifier for the object this phone references to
	TargetID types.UID

	CountryCode string
	AreaCode    string
	Number      string
}
