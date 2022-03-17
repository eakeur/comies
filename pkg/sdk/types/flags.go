package types

type (
	// ProductData is a flag type to point out which product related data an operation must consider
	ProductData int

	WritingFlag int
)

const (

	// Stock points out that stock information properties should be retrieved
	Stock ProductData = iota

	// Sale points out that sale information properties should be retrieved
	Sale ProductData = iota

	// All points out that all its information properties should be retrieved
	All ProductData = iota

	// Overwrite flags out that if a resource exists, the API must overwrite it with the incoming one
	Overwrite WritingFlag = iota

	// Copy flags out that if a resource exists, the API must copy it in another address
	Copy WritingFlag = iota
)
