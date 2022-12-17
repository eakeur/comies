import { Accordion, AccordionItem, Flex, Text } from "@chakra-ui/react";
import { listProducts } from "api/menu";
import { Product } from "core/product";
import { useEffect } from "react";
import { Control, useWatch } from "react-hook-form";
import { useQuery } from "react-query";

interface Props {
    searchControl: Control,
    onSelected: (p: Product) => any
}


export function OrderItemSearcher({searchControl, onSelected}: Props){

    const search = useWatch({control: searchControl, name: "product_name", defaultValue: ""})

    const {data:products, refetch} = useQuery('order-item-products', () => {
        return listProducts({
            types: [Product.outputType, Product.outputCompositeType],
            code: search,
            name: search
        })
    }, {enabled: false})

    useEffect(() => {
        refetch()
    }, [search, refetch])

    return (
        <Accordion defaultIndex={[0]} allowMultiple>
            {
                products?.map(p => {
                    return (
                        <AccordionItem 
                            key={p.id} 
                            display="flex" 
                            flexDirection="row"
                            tabIndex={0}
                            justifyContent="space-between"
                            onClick={() => onSelected(p)}
                            >
                            <Text width="90%">{p.name}</Text>
                            <Flex width="10%" justifyContent="left">
                                <Text fontSize="large">R$ {p.sale_price?.toFixed(2)}</Text>
                            </Flex>
                        </AccordionItem>
                    )
                })
            }
        </Accordion>
    )
}