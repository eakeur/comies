package store

import (
	"gomies/app/core/entities/contacting"
	"gomies/app/core/entities/preferences"
	"gomies/app/core/types/document"
	"gomies/app/core/types/entity"
	"gomies/app/core/types/id"
)

type Store struct {
	entity.Entity
	Name        string
	Nick        string
	Document    document.Document
	Responsible string
	Addresses   []contacting.Address
	Phones      []contacting.Phone
	Preferences preferences.Preferences
}

func (s Store) IsPrimary() bool {
	return s.StoreID == 0 && (s.StoreExternalID == id.Nil || s.StoreExternalID == id.External{})
}
