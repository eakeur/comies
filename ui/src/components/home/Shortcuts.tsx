import { Flex, Input, Text } from "@chakra-ui/react";
import { runAction } from "../../handlers/handler";
import { HomeShortcutButton } from "./ShortcutButton";

export function HomeShortcuts() {  
    return (
        <Flex wrap="wrap" direction="column" >
            <Text as={"h1"} fontSize={"4xl"} mb="10px" fontWeight="bold">Comies</Text>

            
            <Input
                type={'text'}
                variant="filled"
                placeholder='Busque tudo...'
            />


            <Text as={"h4"} fontSize={"md"} mt="30px">O que vamos fazer hoje?</Text>
            <HomeShortcutButton 
                name="Novo pedido" 
                shortcut="Shift + A" 
                icon="BuildQueueNew" 
                color="#ef6c00"
                onClick={() => runAction("add_order")}/>
            <HomeShortcutButton 
                name="Novo produto" 
                shortcut="Shift + D" 
                icon="AppIconDefaultAdd" 
                color="#1565c0"
                onClick={() => runAction("add_product")}/>
            <HomeShortcutButton 
                name="Verificar estoque" 
                shortcut="Shift + V" 
                icon="ProductionFloorManagement" 
                color="#c62828"
                onClick={() => runAction("verify_stock")}/>
            <HomeShortcutButton 
                name="Nova entrada/saída" 
                shortcut="Shift + M" 
                icon="SwitcherStartEnd" 
                color="#6a1b9a"
                onClick={() => runAction("add_stock_movement")}/>
        </Flex>
    )
}