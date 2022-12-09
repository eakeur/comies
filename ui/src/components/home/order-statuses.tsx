import { Flex, SkeletonText, Text } from "@chakra-ui/react";
import { darken } from "polished";
import { useQuery } from "react-query";
import { Icon } from "../shared/Icon";
import { getOrderCountByStatus } from "../../api/comies/ordering";
import { motion } from "framer-motion";
import { ListShowUpAnimation } from "../animations/animations";
import { ErrorBanner } from "../shared/error";
import { Order } from "../../core/order";
import styled from "styled-components";

const statuses = [
  Order.pendingStatus,
  Order.preparingStatus,
  Order.waitingTakeoutStatus,
  Order.waitingDeliveryStatus,
  Order.deliveringStatus,
];

export function OrderStatuses() {
  const { isLoading, isError, data, refetch } = useQuery(
    "order-count-by-status",
    getOrderCountByStatus
  );

  if (isError) {
    return (
      <ErrorBanner
        description="Não foi possível obter o número de pedidos por status"
        retry={refetch}
      />
    );
  }

  return (
    <MotionFlex
      initial="hidden"
      animate="visible"
      variants={ListShowUpAnimation.parent}
      wrap="wrap"
    >
      {statuses.map((s) => {
        const { icon, color, name } = Order.StatusData[s];
        const count = Order.getStatusCount(data, s);

        return (
          <StyledOrderStatusBlock
            variants={ListShowUpAnimation.children}
            key={s}
            color={color}
            tabIndex={1}
            direction="column"
            justify="space-between"
          >
            <Icon name={icon} />
            {isLoading ? (
              <SkeletonText
                endColor="white.200"
                mt="4"
                noOfLines={2}
                spacing="1"
                skeletonHeight="1"
              />
            ) : (
              <h4>{count}</h4>
            )}
            <Text fontSize="sm">{name}</Text>
          </StyledOrderStatusBlock>
        );
      })}
    </MotionFlex>
  );
}

const MotionFlex = motion(Flex);

const StyledOrderStatusBlock = styled(motion(Flex))`
  width: 8rem;
  height: 8rem;
  padding: 15px;
  margin: 5px;
  border-radius: 7px;
  box-shadow: rgba(0, 0, 0, 0.24) 0px 3px 8px;
  background: ${(props) => props.color};

  * {
    color: ${(props) => darken(0.35, props.color ?? "#fff")};
    text-align: left;
    font-weight: bold;
  }

  i {
    font-size: 18px;
  }

  h4 {
    font-size: 24px;
  }
`;
