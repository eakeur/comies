import { Product } from "core/product";
import API from "api";
import axios from 'axios'

export function addProduct(prod: Product) {
    return Promise.resolve({
        ...prod,
        sale_price: API.toCents(prod.sale_price),
        cost_price: API.toCents(prod.cost_price)
    })
    .then((prod) => axios.post(API.menu.addProduct.toString(), prod))
    .then((res) => res.headers[API.idHeader] ?? "")
}

export function updateProduct(productID: string, prod: Product) {
    return Promise.resolve({
        ...prod,
        sale_price: API.toCents(prod.sale_price),
        cost_price: API.toCents(prod.cost_price)
    })
    .then((prod) => axios.post(
        API.menu.addProduct
            .params("product_id", productID)
            .toString(), 
        prod
    ))
    .then((res) => res.headers[API.idHeader] ?? "")
}


export function listProducts(filter?: Product.Filter) {
    return axios.get<Product[]>(
        API.menu.addProduct
            .query(filter)
            .toString(),
    )
    .then((res) => res.data.map(p => ({
        ...p,
        sale_price: API.fromCents(p.sale_price),
        cost_price: API.fromCents(p.cost_price)
    } as Product)))
}