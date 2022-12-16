package types

type Currency int64

type Discount float64

type Amount struct {
	Value     Currency
	Discounts Currency
	Currency  string
	Net       Currency
}
