import { fail } from "core/failures";

namespace API {
    export const idHeader = "Location"

    export function authorize(init: RequestInit = {}): RequestInit {
        return init;
    }

    export async function handle(res: globalThis.Response) {
        if (!res.ok) {
            throw res;
        }

        return res
    }

    export function toCents(value?: number) {
        return (value ?? 0) * 100
    }

    export function fromCents(value?: number) {
        return (value ?? 0) / 100
    }

    export function json(res: globalThis.Response) {
        return res.json().catch(fail);
    }

    class URL extends String {
        params(key: string, value: string): URL {
            return new URL(this.replace(`{${key}}`, value));
        }

        query(query: any): URL {
            return !query ? this : new URL(`${this}?${new URLSearchParams(query).toString()}`)
        }

        request(init?: RequestInit) {
            return fetch(this.toString(), authorize(init)).then(handle)
        }
    }

    export function route(name: string) {
        return new URL(name);
    }

    const menuV1 = process.env.REACT_APP_MENU_V1_URL!
    const orderingV1 = process.env.REACT_APP_ORDERING_V1_URL!

    export const menu = {
        addProduct: new URL(`${menuV1}/items`),
        listProducts: new URL(`${menuV1}/items`),
        updateProduct: new URL(`${menuV1}/items/{product_id}`),
        getProductByID: new URL(`${menuV1}/items/{product_id}`),
        getProductNameByID: new URL(`${menuV1}/items/{product_id}`),

        addIngredient: new URL(`${menuV1}/items/{product_id}/ingredients`),
        listIngredients: new URL(`${menuV1}/items/{product_id}/ingredients`),
        removeIngredient: new URL(`${menuV1}/items/{product_id}/ingredients/{ingredient_id}`),

        addMovement: new URL(`${menuV1}/items/{product_id}/movements`),
        listMovements: new URL(`${menuV1}/items/{product_id}/movements`),
        getMovementsBalance: new URL(`${menuV1}/items/{product_id}/movements/balance`),
        removeMovement: new URL(`${menuV1}/items/{product_id}/movements/{movement_id}`),

        addPrice: new URL(`${menuV1}/items/{product_id}/prices`),
        listPrices: new URL(`${menuV1}/items/{product_id}/prices`),
        getLatest: new URL(`${menuV1}/items/{product_id}/prices/latest`),
    } as const

    export const ordering = {
        getOrderCountByStatus: new URL(`${orderingV1}/orders?statuscount={status}`),
        getOrderStatusByCustomerPhone: new URL(`${orderingV1}/orders/{customer_phone}?phone=true`),
        listSaleableItems: new URL(`${menuV1}/items/saleable`),
    } as const
}

export default API;