package types

import "strconv"

const QuantityMinimum Quantity = 0

type Quantity int64

func (q Quantity) String() string {
	return strconv.FormatInt(int64(q), 10)
}
