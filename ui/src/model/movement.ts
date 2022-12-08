export interface Movement {
    id: string;
    product_id: string;
    type: number;
    date: Date;
    quantity: number;
    paid_value: number;
}


export namespace Movement {
    export const noneType = 0
    export const inputType = 10
    export const outputType = 20

}