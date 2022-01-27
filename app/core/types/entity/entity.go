package entity

import (
	"gomies/app/core/types/audit"
	"gomies/app/core/types/id"
)

// Entity is a type that wraps identifiers and auditioning properties
type Entity struct {
	//ID is the unique identifier of this entity, only known by the boundaries of the domain
	ID id.ID

	// ExternalID is the unique identifier of this entity, known by the outer world
	ExternalID id.External

	audit.Audit

	// Active points out if the entity is being recognised. It can have different side effects according to the entity
	Active bool
}
