package wrappers

import (
	"gomies/app/core/managers/session"
	"gomies/app/core/managers/transaction"
)

type Managers struct {
	Authorization session.Manager
	Transaction   transaction.Manager
}
