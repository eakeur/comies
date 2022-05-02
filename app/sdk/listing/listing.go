package listing

type (
	SortDirection string

	Filter struct {
		SortBy        string
		RangeStart    int
		RangeEnd      int
		SortDirection SortDirection
	}
)

func HasMore(len, cap int) bool {
	return len < cap
}
