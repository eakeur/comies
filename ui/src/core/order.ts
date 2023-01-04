export namespace Ordering {

  export enum OrderStatus {
    pending = 1,
    preparing = 2,
    waitingTakeout = 3,
    waitingDelivery = 4,
    delivering = 5,
  }

  export enum DeliveryType {
    takeout = 10,
    delivery = 20,
  }

  export type Order = {
    id: string;
    placed_at: Date;
    customer_name: string;
    customer_address: string;
    customer_phone: string;
    delivery_type: DeliveryType;
    observations: string;
  }

  export type Ticket = {
    customer_name: string;
    customer_address: string;
    customer_phone: string;
    delivery_type: DeliveryType;
    observations: string;
  }

  export type SaleableItem = {
    id: string,
    name: string,
    code: string,
    price: number,
    stock: number,
  }

  export type TicketItem = {
    saleable: SaleableItem,
    quantity: number;
    observations: string;
  }

  export type CountByStatus = {
    count: number
  };

  export type CurrentStatus = {
    value: number;
    occurred_at: Date;
    order_id: string;
    customer_name?: string;
  };

  export const minuteDifference = (occurence: Date, base = new Date()) =>
    ((base.getTime() - occurence.getTime()) / 60000).toFixed(0);


  export const saleableItemToTicketItem = (saleable: SaleableItem) => ({
    quantity: 1,
    observations: "",
    saleable,
  })
}
