import { useToast } from "@chakra-ui/react";
import { as } from "core/failures";
import { useForm } from "react-hook-form";
import { useMutation } from "react-query";
import { addProduct, updateProduct } from "api/menu";
import { Product } from "core/product";


export function useProductForm(product?: Product) {
    const toast = useToast()
    const control = useForm({ defaultValues: product })

    const mutation = useMutation((product: Product) => {
        return (
            product.id ?
                updateProduct(product.id, product).then((_) => product.id) :
                addProduct(product))
    }, {
        onError(error) {
            const fail = as(error)
            toast({
                title: fail.title,
                status: "error",
                description: fail.description,
            })
        },

        onSuccess(data) {
            toast({
                title: "Produto criado com sucesso",
                status: "success",
            })
        },
    });

    const submit = control
        .handleSubmit((product: Product) =>
            mutation.mutate(Product.clear(product)))

    return { product, control, submit }
}