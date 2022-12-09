import { Product } from "core/product";
import { API } from "./api";

export function addProduct(prod: Product) {
    return API.menu.addProduct
        .request({ method: "POST", body: JSON.stringify(prod) })
        .then((res) => res.headers.get(API.idHeader) ?? "")
}

export function updateProduct(productID: string, prod: Product) {
    return API.menu.updateProduct
        .params("product_id", productID)
        .request({ method: "PUT", body: JSON.stringify(prod) })
}
