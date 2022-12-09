import { Order } from "../core/order"
import { API } from "api/api"
import { Routes } from "api/comies/routes"

export function getOrderCountByStatus(status?: number) {
    return API
        .route(Routes.ordering.getOrderCountByStatus)
        .params("status", status?.toString() ?? '')
        .request()
        .then((res) => res.json())
        .then((res) => Number.parseInt(res.count))
}

export function getOrderStatusByCustomerPhone(phone: string) {
    return API
        .route(Routes.ordering.getOrderStatusByCustomerPhone)
        .params("customer_phone", phone)
        .request()
        .then((res) => res.json())
        .then<Order.CurrentStatus>((res) => ({
            ...res,
            occurred_at: new Date(res.occurred_at)
        }))
}