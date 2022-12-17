import { Box, Text, Flex, FormControl, FormLabel, Input, Stack, Accordion, AccordionItem } from "@chakra-ui/react";
import { listProducts } from "api/menu";
import { Order } from "core/order";
import { Product } from "core/product";
import { useEffect, useState } from "react";
import { useQuery } from "react-query";
import { useDebounce } from "use-debounce";

interface ItemSearchProps {
    onSelect(item: Order.Item): void
}

export function ItemSearch({onSelect}: ItemSearchProps){

    const [search, setSearch] = useState("");

    const [term] = useDebounce(search, 500);
    
    const {data, refetch} = useQuery('order-item-products', () => {
        return listProducts({
            code: search,
            name: search,
            types: [
                Product.outputType, 
                Product.outputCompositeType,
            ],
        })
    }, {enabled: false})

    useEffect(() => {
        if (term && term.length >= 3) refetch()
    }, [term, refetch])

    function select(p: Product){
        onSelect({
            quantity: 1,
            discounts: 0,
            observations: "",
            product: {
                id: p.id!,
                name: p.name!,
                code: p.code!,
                price: p.sale_price!,
            },
        })
    }
    
    return (
        <Box>
            <FormControl>
                <FormLabel fontSize="sm">Nome</FormLabel>
                <Input defaultValue={search} onChange={(ev) => setSearch(ev.target.value)}/>
            </FormControl>
            
            <Accordion defaultIndex={[0]} allowMultiple>
            {
                data?.map(p => {
                    return (
                        <AccordionItem 
                            key={p.id} 
                            display="flex" 
                            flexDirection="row"
                            tabIndex={0}
                            justifyContent="space-between"
                            paddingBlock="4px"
                            onClick={() => select(p)}
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
        </Box>
    )
}
