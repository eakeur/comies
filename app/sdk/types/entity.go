package types

import (
	"time"
)

type (
	Entity struct {
		//ID is the unique identifier of this entity, only known by the boundaries of the domain
		ID ID
		// Active points out if the entity is being recognised. It can have different side effects according to the entity
		Active bool
		// Every time an entity enters a context that will change it, a history entry must be set here
		History
	}

	Store struct {
		// StoreID is the ID of the store this entity belongs to
		StoreID ID
	}

	History struct {
		By        ID
		At        time.Time
		Operation string
	}
)
