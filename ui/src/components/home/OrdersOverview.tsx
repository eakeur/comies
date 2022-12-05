import { ArrowForwardIcon } from "@chakra-ui/icons";
import { InputGroup, InputRightElement, Input, Button, Text, Flex } from "@chakra-ui/react";

import styled from "styled-components"
import { OrderStatusBlock } from "./OrderStatusBox"

const StyledOrdersOverview = styled.section`
    > * {
        margin-bottom: 30px;
    }
`

export function OrdersOverview(){
    return (
        <StyledOrdersOverview>
            <Text as={"h1"} fontSize={"2xl"}>Meus pedidos</Text>

            <Button width="100%" rightIcon={<ArrowForwardIcon />} colorScheme='orange' variant='outline'>
                Ver painel de pedidos
            </Button>

            <Text as={"label"} fontSize={"md"}>Verifique o status de um pedido:</Text>
            <InputGroup size='md'>
                <Input
                    pr='4.5rem'
                    type={'phone'}
                    placeholder='Telefone do cliente'
                    variant="filled"
                />
                <InputRightElement width='7rem'>
                    <Button h='1.75rem' size='sm' aria-label="Pesquisar pedido por telefone do cliente">
                        Pesquisar
                    </Button>
                </InputRightElement>
            </InputGroup>

            <Text as={"h3"} fontSize={"xl"}>Resumo do dia</Text>
            <Flex wrap="wrap">
                <OrderStatusBlock status={1} quantity={3}/>
                <OrderStatusBlock status={2} quantity={7}/>
                <OrderStatusBlock status={3} quantity={2}/>
                <OrderStatusBlock status={4} quantity={1}/>
                <OrderStatusBlock status={5} quantity={5}/>
            </Flex>
        </StyledOrdersOverview>
    )
}