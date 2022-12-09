
const menuV1 = process.env.REACT_APP_MENU_V1_URL!
const orderingV1 = process.env.REACT_APP_ORDERING_V1_URL!

export namespace Routes {
    export const menu = {
        addProduct: `${menuV1}/products`,
        listProducts: `${menuV1}/products`,
        updateProduct: `${menuV1}/products/{product_id}`,
        getProductByID: `${menuV1}/products/{product_id}`,
        getProductNameByID: `${menuV1}/products/{product_id}`,

        addIngredient: `${menuV1}/products/{product_id}/ingredients`,
        listIngredients: `${menuV1}/products/{product_id}/ingredients`,
        removeIngredient: `${menuV1}/products/{product_id}/ingredients/{ingredient_id}`,

        addMovement: `${menuV1}/products/{product_id}/movements`,
        listMovements: `${menuV1}/products/{product_id}/movements`,
        getMovementsBalance: `${menuV1}/products/{product_id}/movements/balance`,
        removeMovement: `${menuV1}/products/{product_id}/movements/{movement_id}`,

        addPrice: `${menuV1}/products/{product_id}/prices`,
        listPrices: `${menuV1}/products/{product_id}/prices`,
        getLatest: `${menuV1}/products/{product_id}/prices/latest`,
    } as const
    
    export const ordering = {
        getOrderCountByStatus: `${orderingV1}/orders?statuscount={status}`,
        getOrderStatusByCustomerPhone: `${orderingV1}/orders/{customer_phone}?phone=true`,
    }
}