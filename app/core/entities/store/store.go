package store

import (
	"gomies/app/sdk/types"
)

type Store struct {
	types.Entity
	Name            string
	Nickname        string
	ResponsibleName string
	Preferences     types.Preferences
	Document        types.Document
}
