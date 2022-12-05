import { Ingredient } from "../model/ingredient";
import { Movement } from "../model/movement";
import { Product } from "../model/product";

const route = `${process.env.REACT_APP_API_URL}`

const headers = {}

interface APIResponse {
    data?: any,
    error?: {
        message: string,
        code: string,
    }
}

export namespace Comies {
    export function saveProduct(prod: Product) {
        if (prod.id) {
            return fetch(`${route}/menu/products/${prod.id}`, {method: "PUT", headers: headers, body: JSON.stringify(prod)})
                .then<void>(handle())
        }

        return fetch(`${route}/menu/products`, {method: "POST", headers: headers, body: JSON.stringify(prod)})
        .then<{id: string}>(handle()).then(payload => payload.id)
    }

    export function listProducts(filter: any) {
        return fetch(`${route}/menu/products${query(filter)}`, {headers: headers, body: JSON.stringify(filter)})
    }

    export function getProductByKey(key: string) {
        return fetch(`${route}/menu/products/${key}`, {headers: headers}).then<Product>(handle())
    }
    
    export function getProductProperty(id: string, prop: 'name' | 'stock-balance') {
        return fetch(`${route}/menu/products/${id}/${prop}`, {headers: headers})
        .then<{balance: number}>(handle()).then(p => p.balance)
    }

    export function removeProduct(id: string){
        return fetch(`${route}/menu/products/${id}`, {method: "DELETE", headers: headers})
        .then<void>(handle())
    }

    export function saveIngredient(ingredient: Ingredient){
        return fetch(`${route}/menu/products/${ingredient.product_id}/ingredients`, {
            method: "POST", headers: headers, body: JSON.stringify(ingredient),
        }).then<{id: string}>(handle())
    }

    export function listIngredients(productID: string){
        return fetch(`${route}/menu/products/${productID}/ingredients`, {headers: headers})
        .then<Ingredient[]>(handle())
    }

    export function removeIngredient(productID: string, id: string){
        return fetch(`${route}/menu/products/${productID}/ingredients/${id}`, {method: "DELETE", headers: headers})
        .then<void>(handle())
    }

    export function saveMovement(movement: Movement){
        return fetch(`${route}/menu/products/${movement.product_id}/movements`, {
            method: "POST", headers: headers, body: JSON.stringify(movement)
        }).then<{id: string}>(handle())
    }

    export function listMovements(productID: string){
        return fetch(`${route}/menu/products/${productID}/movements`, {headers: headers})
        .then<Movement[]>(handle())
    }

    export function removeMovement(productID: string, id: string){
        return fetch(`${route}/menu/products/${productID}/movements/${id}`, {method: "DELETE", headers: headers})
        .then<void>(handle())
    }
}

function query(filters: any){
    if (filters) {
        return "?" + new URLSearchParams(filters).toString()
    }

    return ""
}

function handle(parser?: (data: any) => any) {
    return async function(res: globalThis.Response) {
        if (res.status === 204) {
            return 
        }
    
        let payload: APIResponse;
    
        try {
            payload = await res.json();
        } catch (error) {
            payload = { error: {code: "", message: ""}}
        }
    
        if (payload?.error) {
            throw payload.error
        }
    
        return parser?.call(undefined, payload?.data) ?? payload.data
    }
}