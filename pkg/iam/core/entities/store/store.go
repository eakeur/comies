package store

import (
	"gomies/pkg/contacting/core/entities/contacting"
	"gomies/pkg/sdk/types"
)

type Store struct {
	types.Entity
	Name             string
	Nickname         string
	ResponsibleName  string
	ParentID         types.ID
	ParentExternalID types.External
	Preferences      types.Preferences
	Document         types.Document
	Addresses        []contacting.Address
	Phones           []contacting.Phone
}

func (s Store) IsPrimary() bool {
	return s.ParentID == 0 && s.ParentExternalID.Empty()
}
