package history

import (
	"gomies/app/core/types/id"
	"time"
)

// History wraps information about entity creation and changes
type History struct {
	By        id.External
	At        time.Time
	Operation string
}
