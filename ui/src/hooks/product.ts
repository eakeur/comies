import { useEffect, useState } from "react";
import { useForm } from "react-hook-form";
import { Product, ProductType, Unit } from "../model/product";
import { Comies }  from "../services/comies";

export function useProduct(id: string) {
    const [product, setProduct] = useState<Product>({});

    useEffect(function () {
        Comies.getProductByKey(id).then(setProduct)
    }, [id])

    return product;
}

export function useProductForm(product?: Product) {
    const control = useForm({defaultValues: product})

    const submit = control.handleSubmit(function(product: Product) {
        product = Product.analyseAndClean(product)
        Comies?.saveProduct(product)
    })

    return {product, control, submit}
}