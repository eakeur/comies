import { Order } from "../model/order"
import { API } from "./api"

export function getOrderCountByStatus() {
    return fetch(API.route("GET_COUNT_BY_STATUS_URL"), API.authorize())
        .then(API.response)
        .then<Order.CountByStatus>((res) => res.json())

}

export function getOrderStatusByCustomerPhone(phone: string) {
    return fetch(
        API.route("GET_STATUS_BY_CUSTOMER_PHONE_URL", {params: {"{customer_phone}": phone}}),
        API.authorize()
    )
        .then(API.response)
        .then<Order.CurrentStatus>((res) => res.json())
        .then<Order.CurrentStatus>((res) => ({
            ...res,
            occurred_at: new Date(res.occurred_at)
        }))
}