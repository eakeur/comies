import { ArrowForwardIcon } from "@chakra-ui/icons";
import {
    Button,
    Flex,
    Text,
} from "@chakra-ui/react";
import { ListShowUpAnimation } from "components/animations/animations";
import { Ordering } from "core/order";
import { motion } from "framer-motion";
import { OrderStatusChecker } from "./status-checker";
import { OrderCountByStatus } from "./order-status";

export function MyOrders() {

    return (
        <section id="myorders">
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
                    Checar pedido do cliente:
                </Text>
                <OrderStatusChecker/>
            </div>

            <div>
                <Text as={"h3"} fontSize={"l"}>
                    Resumo do dia
                </Text>
                <StatusOverview
                    variants={ListShowUpAnimation.parent}
                    wrap="wrap"
                    initial="hidden"
                    animate="visible">
                    {statuses.map((s) => (
                        <OrderCountByStatus
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
    Ordering.OrderStatus.pending,
    Ordering.OrderStatus.preparing,
    Ordering.OrderStatus.waitingTakeout,
    Ordering.OrderStatus.waitingDelivery,
    Ordering.OrderStatus.delivering,
];

const StatusOverview = motion(Flex);
