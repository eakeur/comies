package store

import (
	"gomies/app/core/entities/contacting"
	"gomies/app/sdk/types"
)

type Store struct {
	types.Entity
	Name            string
	Nickname        string
	ResponsibleName string
	ParentID        types.ID
	Preferences     types.Preferences
	Document        types.Document
	Addresses       []contacting.Address
	Phones          []contacting.Phone
}

func (s Store) IsPrimary() bool {
	return s.ParentID.Empty()
}
