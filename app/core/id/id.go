package id

import (
	"fmt"
	"strconv"
	"strings"
)

type ID int64

func (i ID) String() string {
	return strconv.FormatInt(int64(i), 10)
}

func (i ID) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, i.String())), nil
}

func (i *ID) UnmarshalJSON(b []byte) error {
	id, err := FromString(strings.ReplaceAll(string(b), `"`, ""))
	if err != nil {
		return err
	}

	*i = id

	return nil
}

func FromString(in string) (ID, error) {
	i, err := strconv.Atoi(in)
	if err != nil {
		return 0, err
	}

	return ID(i), nil
}

func ValidateID(id ID) error {
	if id <= 0 {
		return ErrNoID
	}

	return nil
}
