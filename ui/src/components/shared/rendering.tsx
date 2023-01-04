import { Ordering } from "core/order";

export const OrderStatusRenderingData: {
    readonly [key: number]: { name: string; icon: string; color: string };
  } = {
    [Ordering.OrderStatus.pending]: {
      icon: "Clock",
      color: "#90a4ae",
      name: "Pendente",
    },
    [Ordering.OrderStatus.preparing]: {
      icon: "EatDrink",
      color: "#ffb74d",
      name: "Preparando",
    },
    [Ordering.OrderStatus.waitingTakeout]: {
      icon: "OfficeStoreLogo",
      color: "#4dd0e1",
      name: "Pra retirar",
    },
    [Ordering.OrderStatus.waitingDelivery]: {
      icon: "Assign",
      color: "#e57373",
      name: "Pra entregar",
    },
    [Ordering.OrderStatus.delivering]: {
      icon: "DeliveryTruck",
      color: "#81c784",
      name: "Entregando",
    },
  } as const;