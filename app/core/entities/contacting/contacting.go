package contacting

import "gomies/app/sdk/types"

type (
	Address struct {
		types.Entity
		TargetID   types.ID
		Code       string
		Street     string
		Number     string
		Complement string
		District   string
		City       string
		State      string
		Country    string
	}

	Phone struct {
		types.Entity
		TargetID    types.ID
		CountryCode string
		AreaCode    string
		Number      string
	}

	Contact struct {
		types.Entity
		TargetID  types.ID
		Addresses []Address
		Phones    []Phone
	}
)

func (c *Contact) AddAddresses(addresses ...Address) {
	c.Addresses = append(c.Addresses, addresses...)
}

func (c *Contact) AddPhones(phones ...Phone) {
	c.Phones = append(c.Phones, phones...)
}
