import { Product } from "../model/product";
import { API } from "./api";
import { Routes } from "./comies/routes";

export function addProduct(prod: Product) {
    return API.route(Routes.menu.addProduct)
        .request({ method: "POST", body: JSON.stringify(prod) })
        .then((res) => res.headers.get(API.idHeader) ?? "")
}

export function updateProduct(productID: string, prod: Product) {
    return API.route(Routes.menu.updateProduct)
        .params("product_id", productID)
        .request({ method: "PUT", body: JSON.stringify(prod) })
}
