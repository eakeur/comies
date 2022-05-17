package store

import (
	"gomies/app/core/entities/address"
	"gomies/app/core/entities/phone"
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
	Addresses       []address.Address
	Phones          []phone.Phone
}

func (s Store) IsPrimary() bool {
	return s.ParentID.Empty()
}
