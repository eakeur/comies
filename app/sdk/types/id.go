package types

type (
	ID int64
)

func (e ID) Empty() bool {
	return e == 0
}
