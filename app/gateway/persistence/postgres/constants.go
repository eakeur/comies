package postgres

const (
	DuplicateError = "23505"
	NonexistentFK  = "23503"

	ItemOrderIDFK = "order_id_fk"
	ItemIDPK      = "items_pk"

	IngredientProductUK   = "product_ingredient_id_uk"
	IngredientProductIDFK = "product_id_fk"
	IngredientIDFK        = "ingredient_id_fk"
)
