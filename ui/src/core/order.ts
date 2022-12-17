export interface Order {}

export namespace Order {
  export const pendingStatus = 1;
  export const preparingStatus = 2;
  export const waitingTakeoutStatus = 3;
  export const waitingDeliveryStatus = 4;
  export const deliveringStatus = 5;

  export const statuses = [
    pendingStatus,
    preparingStatus,
    waitingTakeoutStatus,
    waitingDeliveryStatus,
    deliveringStatus,
  ] as const;

  export const StatusData: {
    readonly [key: number]: { name: string; icon: string; color: string };
  } = {
    [pendingStatus]: {
      icon: "Clock",
      color: "#90a4ae",
      name: "Pendente",
    },
    [preparingStatus]: {
      icon: "EatDrink",
      color: "#ffb74d",
      name: "Preparando",
    },
    [waitingTakeoutStatus]: {
      icon: "OfficeStoreLogo",
      color: "#4dd0e1",
      name: "Pra retirar",
    },
    [waitingDeliveryStatus]: {
      icon: "Assign",
      color: "#e57373",
      name: "Pra entregar",
    },
    [deliveringStatus]: {
      icon: "DeliveryTruck",
      color: "#81c784",
      name: "Entregando",
    },
  } as const;

  export const minuteDifference = (occurence: Date, base = new Date()) =>
    ((base.getTime() - occurence.getTime()) / 60000).toFixed(0);

  export type CountByStatus = {count: number};
  export type CurrentStatus = {
    value: number;
    occurred_at: Date;
    order_id: string;
    customer_name?: string;
  };
}
