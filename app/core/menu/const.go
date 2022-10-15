package menu

type Type int

const NoType Type = 0

const (
	ReplaceIngredientChangeType  Type = 1
	IgnoreIngredientChangeType   Type = 2
)

const (
	OutputProductType          Type = 10
	OutputCompositeProductType Type = 20
	InputProductType           Type = 30
	InputCompositeProductType  Type = 40
)


const (
	InputMovementType    Type = 10
	OutputMovementType   Type = 20
	ReservedMovementType Type = 30
)