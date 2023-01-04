import { listSaleableItems } from "api/ordering";
import { Ordering } from "core/order";
import { useState, useEffect, createContext } from "react";
import { useForm } from "react-hook-form";
import { useQuery } from "react-query";
import { useDebounce } from "use-debounce";

export const ItemsContext = createContext<Ordering.TicketItem[]>([]);

export function useOrderPlacement(){
    const items = useTicketItems()

    const form = useForm()

    const search = useSaleableItemsSearch()

    return {
        form: form,
        search,
        ...items,
    }
}

export function useTicketItems(){
    const [ items, setItems ] = useState<Ordering.TicketItem[]>([]);
    return {
        list: items, setItems,

        addItem(it: Ordering.TicketItem){
            setItems(items.concat(it))
        },

        removeItem(idx: number){
            setItems(items.filter((_, i) => i !== idx))
        },

        setItem(it: Ordering.TicketItem, index: number){
            setItems(items.map((item, idx) => idx === index ? it : item))
        }
    }
}

export function useSaleableItemsSearch() {

    const [search, setSearch] = useState("");

    const [term] = useDebounce(search, 250);

    const { data, refetch, isLoading, isError } = useQuery('order-item-products', () => listSaleableItems(term), {
        enabled: false,
    })

    useEffect(() => {
        if (term && term.length >= 3) refetch()
    }, [term, refetch])


    return {
        data, isLoading, isError, search, setSearch
    }
}