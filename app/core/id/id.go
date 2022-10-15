package id

import (
	"fmt"
	"strconv"
)

type ID int64

func (i ID) String() string {
	return strconv.FormatInt(int64(i), 10)
}

func (i ID) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, i.String())), nil
}

func ValidateID(id ID) error {
	if id <= 0 {
		return ErrNoID
	}

	return nil
}
