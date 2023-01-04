import {
    Box,
    useDisclosure

} from "@chakra-ui/react";
import { Ordering } from "core/order";
import { SaleableItem, TicketItemEditorModal } from "./saleable-item";

export const SearchResult = ({ saleable, onAdd }: { saleable: Ordering.SaleableItem, onAdd: (item: Ordering.TicketItem) => void }) => {
    const disclosure = useDisclosure()

    const item = Ordering.saleableItemToTicketItem(saleable)

    return <>
        <Box display="grid"
            gridTemplateColumns="1fr auto"
            tabIndex={0}
            onClick={disclosure.onOpen}
            onKeyDown={(e) => e.key === " " || e.code === "Space" ? disclosure.onOpen() : undefined}>
            <SaleableItem saleable={saleable} showStock />
        </Box>
        <TicketItemEditorModal item={item} disclosure={disclosure} add={onAdd} />
    </>
}