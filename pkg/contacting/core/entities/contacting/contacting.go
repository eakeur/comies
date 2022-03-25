package contacting

import "gomies/pkg/sdk/types"

type (
	Address struct {
		types.Entity
		TargetID   types.UID
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
		TargetID    types.UID
		CountryCode string
		AreaCode    string
		Number      string
	}

	Contact struct {
		types.Entity
		TargetID  types.UID
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
