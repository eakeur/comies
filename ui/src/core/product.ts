export interface Product {
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
}

export namespace Product {
  export const noneType = 0;
  export const outputType = 10;
  export const outputCompositeType = 20;
  export const inputType = 30;
  export const inputCompositeType = 40;

  export const milligram = "mg";
  export const kilogram = "kg";
  export const gram = "g";
  export const unit = "un";

  export const statuses = [
    noneType,
    outputType,
    outputCompositeType,
    inputType,
    inputCompositeType,
  ] as const;

  export const ProductTypeData: { readonly [key: number]: { name: string } } = {
    [noneType]: { name: "Nenhum" },
    [outputType]: { name: "Saída" },
    [outputCompositeType]: { name: "Saída composto" },
    [inputType]: { name: "Entrada" },
    [inputCompositeType]: { name: "Entrada composto" },
  } as const;

  export const ProductUnitData: { readonly [key: string]: { name: string } } = {
    [milligram]: { name: "miligramas" },
    [kilogram]: { name: "quilogramas" },
    [gram]: { name: "grama" },
    [unit]: { name: "unidade" },
  } as const;

  export const isComposite = (t?: number) =>
    t === outputCompositeType || t === inputCompositeType;

  export const isOutput = (t?: number) =>
    t === outputType || t === outputCompositeType;

  export const isInput = (t?: number) => t === inputType || t === inputCompositeType;

  /**
   * clear removes some properties from a product object
   * based on its type
   * @param product The product to be cleaned and analysed
   * @returns The cleaned product object
   */
  export function clear(product: Product) {
    if (isComposite(product.type)) {
      product = {
        ...product,
        location: "",
        minimum_quantity: 0,
        maximum_quantity: 0,
      };
    }

    if (isInput(product.type)) {
      product = { ...product, minimum_sale: 0, sale_price: 0 };
    }

    return product;
  }

  export interface Filter {
    types?: number[]
    code?: string
    name?: string
  }
}
