import { AddIcon, EditIcon, ExternalLinkIcon, RepeatIcon } from "@chakra-ui/icons";
import { Accordion, AccordionItem, Input, Menu, MenuButton, MenuItem, MenuList, Text } from "@chakra-ui/react";
import { listProducts } from "api/menu";
import { Product } from "core/product";
import { useEffect, useState } from "react";
import { useQuery } from "react-query";


export function OrderItemSearcher(){

    const [search, setSearch] = useState("")

    const {data:products, refetch, isLoading, isError} = useQuery('order-item-products', () => {
        return listProducts({
            types: [Product.outputType, Product.outputCompositeType],
            code: search,
            name: search
        }).then((p) => {
            console.log(p)
            return p
        })
    }, {enabled: false})

    useEffect(function(){
        if (search.length >= 3) {
            refetch()
        }
    }, [search, refetch])

    return (
        <div>
            
            <Input defaultValue={search} onChange={(ev) => setSearch(ev.target.value)}/>
            {
                search.length >= 3 && 
                <Accordion defaultIndex={[0]} allowMultiple>
                    {
                        products?.map(p => {
                            return <AccordionItem key={p.id}>
                                {p.name}
                                <Text fontSize="large">R${p.sale_price?.toFixed(2)}</Text>
                            </AccordionItem>
                        })
                    }
              </Accordion>
            }
        </div>
    )
}