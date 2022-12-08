import { Product } from "../model/product";
import { API } from "./api";

export function addProduct(prod: Product) {
    return fetch(
        API.route("ADD_PRODUCT_URL"),
        API.authorize({ method: "POST", body: JSON.stringify(prod) })
    )
        .then(API.response)
        .then((res) => res.headers.get(API.idHeader));
}

export function updateProduct(productID: string, prod: Product) {
    return fetch(
        API.route("UPDATE_PRODUCT_URL", {params: {"{product_id}": productID}}),
        API.authorize({ method: "PUT", body: JSON.stringify(prod) })
    ).then(API.response);
}
