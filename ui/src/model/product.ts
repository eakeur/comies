export type Product = {
    id?: string;
    code?: string;
    name?: string;
    type?: number;
    cost_price?: number;
    sale_price?: number;
    sale_unit?: string;
    minimum_sale?: number;
    minimum_quantity?: number;
    maximum_quantity?: number;
    location?: string;
    balance?: string;
}

export const ProductType = {
    none: 0,
    output: 10,
    output_composite: 20,
    input: 30,
    input_composite: 40,
}

export const Unit = {
    milligram: "mg",
    kilogram: "kg",
    gram: "g",
    unit: "un",
}

export namespace Product {
    export function isProductComposite(t: number | undefined){
        return t === ProductType.output_composite || t === ProductType.input_composite
    }
    
    export function isProductOutput(t: number | undefined){
        return t === ProductType.output || t === ProductType.output_composite
    }
    
    export function isProductInput(t: number | undefined){
        return t === ProductType.input || t === ProductType.input_composite
    }

    /**
     * analyseAndClean removes some properties from a product object
     * based on its product type
     * @param product The product to be cleaned and analysed 
     * @returns The cleaned product object
     */
    export function analyseAndClean(product: Product) {
        if (isProductComposite(product.type)){
            product = { ...product,
                location: "",
                minimum_quantity: 0,
                maximum_quantity: 0,
            }
        }

        if (isProductInput(product.type)){
            product = { ...product,
                minimum_sale: 0,
                sale_price: 0,
            }
        }

        return product
    }
}