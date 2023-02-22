import { Ordering } from "core/order"
import API from "api"
import axios from "axios"

export function getOrderCountByStatus(status?: number) {
    return axios.get<Ordering.CountByStatus>(
        API.ordering.getOrderCountByStatus
            .params("status", status?.toString() ?? '')
            .toString()
    ).then((res) => res.data.count)
}

export function getOrderStatusByCustomerPhone(phone: string) {
    return axios.get<Ordering.CurrentStatus>(
        API.ordering.getOrderStatusByCustomerPhone
            .params("customer_phone", phone)
            .toString()
    )
    .then((res) => res.data)
    .then<Ordering.CurrentStatus>((res) => ({
        ...res,
        occurred_at: new Date(res.occurred_at)
    }))
}

export function listSaleableItems(name: string) {
    return axios.get<Ordering.SaleableItem[]>(
        API.menu.listProducts
            .query({identifier: name, saleable: true})
            .toString(),
    ).then(r => r.data)
}