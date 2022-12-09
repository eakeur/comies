import { ArrowForwardIcon } from "@chakra-ui/icons";
import {
    Button,
    Flex,
    Input,
    InputGroup,
    InputRightElement,
    Text,
} from "@chakra-ui/react";
import { ListShowUpAnimation } from "components/animations/animations";
import { Order } from "core/order";
import { motion } from "framer-motion";
import { FormEvent, useState } from "react";
import { OrderStatusChecker } from "./status-checker";
import { OrderCountByStatus } from "./order-status";

export function MyOrders() {
    const [phone, setPhone] = useState("");

    function submitPhone(ev: FormEvent<HTMLFormElement>) {
        const phone = new FormData(ev.currentTarget).get("phone");
        if (phone) setPhone(phone.toString());

        ev.preventDefault();
    }

    return (
        <section>
            <Text as={"h1"} fontSize={"2xl"}>
                Meus pedidos
            </Text>

            <Button
                width="100%"
                marginBottom="30px"
                rightIcon={<ArrowForwardIcon />}
                colorScheme="orange"
                variant="outline"
            >
                Ver painel de pedidos
            </Button>

            <div>
                <Text as={"h3"} fontSize={"l"}>
                    Verifique o status de um pedido:
                </Text>
                <form onSubmit={submitPhone}>
                    <InputGroup size="md">
                        <Input
                            pr="4.5rem"
                            type="phone"
                            name="phone"
                            placeholder="Telefone do cliente"
                            variant="outline"
                            defaultValue={phone}
                        />
                        <InputRightElement width="7rem">
                            <Button
                                type="submit"
                                h="1.75rem"
                                size="sm"
                                aria-label="Pesquisar pedido por telefone do cliente"
                            >
                                Pesquisar
                            </Button>
                        </InputRightElement>
                    </InputGroup>
                    {phone !== "" && <OrderStatusChecker phone={phone} />}
                </form>
            </div>

            <div>
                <Text as={"h3"} fontSize={"l"}>
                    Resumo do dia
                </Text>
                <StatusOverview variants={ListShowUpAnimation.parent}>
                    {statuses.map((s) => (
                        <OrderCounter
                            variants={ListShowUpAnimation.children}
                            key={s}
                            status={s}
                        />
                    ))}
                </StatusOverview>
            </div>
        </section>
    );
}

const statuses = [
    Order.pendingStatus,
    Order.preparingStatus,
    Order.waitingTakeoutStatus,
    Order.waitingDeliveryStatus,
    Order.deliveringStatus,
];

const StatusOverview = motion(Flex);

const OrderCounter = motion(OrderCountByStatus);
