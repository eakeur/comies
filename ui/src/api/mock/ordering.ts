import MockerForNode from "http-request-mock/src/mocker/mocker-for-node";
import { Order } from "../../model/order";

export default function mockOrderingAPIs(mock: MockerForNode){
    const { faker } = require('http-request-mock/http-request-mock.js');

    mock.get(process.env.REACT_APP_GET_COUNT_BY_STATUS_URL!, [
        {status: 1, count: 1},
        {status: 2, count: 5},
        {status: 3, count: 3},
        {status: 4, count: 6},
        {status: 5, count: 4},
    ] as Order.CountByStatus, {
        delay: 2000,
    })
    
    mock.get(new RegExp(process.env.REACT_APP_GET_STATUS_BY_CUSTOMER_PHONE_URL!.replace("{customer_phone}", ".*")), {
        value: faker.rand(1, 5),
        occurred_at: new Date(),
        order_id: "9668257427584",
        customer_name: faker.name(),
    } as Order.CurrentStatus, {
        delay: 1500
    })
}