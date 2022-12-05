import { Flex, Text } from "@chakra-ui/react";
import { darken } from "polished";
import styled from "styled-components";
import { Icon } from "../shared/Icon";

type OrderStatusBlockProps = {
    status: number,
    quantity: number,
}

const OrderStatusData: {status: number, icon: string, color: string, description: string}[] = [
    {status: 1, icon: "Clock", color: "#90a4ae", description: "Pendente"},
    {status: 2, icon: "EatDrink", color: "#ffb74d", description: "Preparando"},
    {status: 3, icon: "OfficeStoreLogo", color: "#4dd0e1", description: "Pra retirar"},
    {status: 4, icon: "Assign", color: "#e57373", description: "Pra entregar"},
    {status: 5, icon: "DeliveryTruck", color: "#81c784", description: "Entregando"}
]

const StyledOrderStatusBlock = styled(Flex)`
    width: 8rem;
    height: 8rem;
    padding: 15px;
    margin: 5px;
    border-radius: 7px;
    box-shadow: rgba(0, 0, 0, 0.24) 0px 3px 8px;
    background: ${(props) => props.color };

    * {
        color: ${(props) => darken(0.35, props.color ?? '#fff')};
        text-align: left;
        font-weight: bold;
    }
    
    i {
        font-size: 18px;
    }

    h4 {
        font-size: 24px;
    }
`



export function OrderStatusBlock({status: type, quantity}: OrderStatusBlockProps){
    const {icon, color, description} = OrderStatusData.find(x => x.status === type)!
    return (
        <StyledOrderStatusBlock color={color} tabIndex={1} direction="column" justify="space-between">
            <Icon name={icon}/>
            <h4>{quantity}</h4>
            <Text fontSize="sm">{description}</Text>
        </StyledOrderStatusBlock>
    )
}