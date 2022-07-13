package postgres

const (
	DuplicateError = "23505"
	NonexistentFK  = "23503"

	ItemOrderIDFK = "order_id_fk"
	ItemIDPK      = "items_pk"

	IngredientProductUK   = "product_ingredient_id_uk"
	IngredientProductIDFK = "product_id_fk"
	IngredientIDFK        = "ingredient_id_fk"

	OrderIDPK = "orders_pk"

	OrderFlowFK   = "order_id_fk"
	OrderFlowPK   = "flow_pk"
	OrderStatusUK = "order_status_uk"

	MemberIDPK       = "member_pk"
	MemberNicknameUK = "nickname_uk"

	MovementIDPK      = "movements_id"
	MovementStockIDFK = "stock_id_fk"
)
