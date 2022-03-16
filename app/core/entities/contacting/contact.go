package contacting

import "gomies/app/core/types/id"

type Contact struct {
	TargetID  id.External
	Phones    []Phone
	Addresses []Address
}
