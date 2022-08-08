package types

import "strconv"

type (
	ID int64
)

func (i ID) String() string {
	return strconv.FormatInt(int64(i), 10)
}

func (i ID) Empty() bool {
	return i == 0
}
