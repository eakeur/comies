package store

import (
	"gomies/app/sdk/types"
)

type Store struct {
	types.Entity
	Name        string
	Nickname    string
	Preferences types.Preferences
}
