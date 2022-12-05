export type Movement = {
    id: string;
    product_id: string;
    type: number;
    date: Date;
    quantity: number;
    paid_value: number;
}

export const MovementType = {
    none: 0,
    input: 10,
    output: 20,
    reserved: 30
}