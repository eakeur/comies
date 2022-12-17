import { Order } from "core/order"
import API from "api"
import axios from "axios"

export function getOrderCountByStatus(status?: number) {
    return axios.get<Order.CountByStatus>(
        API.ordering.getOrderCountByStatus
            .params("status", status?.toString() ?? '')
            .toString()
    ).then((res) => res.data.count)
}

export function getOrderStatusByCustomerPhone(phone: string) {
    return axios.get<Order.CurrentStatus>(
        API.ordering.getOrderStatusByCustomerPhone
            .params("customer_phone", phone)
            .toString()
    )
    .then((res) => res.data)
    .then<Order.CurrentStatus>((res) => ({
        ...res,
        occurred_at: new Date(res.occurred_at)
    }))
}