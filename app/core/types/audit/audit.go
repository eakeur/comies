package audit

import "time"

// Audit wraps information about entity creation and changes
type Audit struct {

	// CreatedAt points out the time the entity was created
	CreatedAt time.Time

	// UpdatedAt points out the time the entity was updated
	UpdatedAt time.Time
}
