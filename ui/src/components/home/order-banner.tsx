import { useQuery } from "react-query";
import { Flex, SkeletonCircle, SkeletonText, Text } from "@chakra-ui/react";
import { darken } from "polished";
import styled from "styled-components";
import { Order } from "../../core/order";
import { getOrderStatusByCustomerPhone } from "../../api/comies/ordering";
import { Icon } from "../shared/Icon";
import { ArrowForwardIcon } from "@chakra-ui/icons";
import { ErrorBanner } from "../shared/error";

export function OrderBanner({ phone }: { phone: string }) {
  const { data, isError, isLoading, refetch } = useQuery(phone, () =>
    getOrderStatusByCustomerPhone(phone)
  );

  if (isLoading) {
    return (
      <StyledFlex color="#ececec">
        <SkeletonCircle size="7" />
        <SkeletonText mt="4" noOfLines={1} spacing="4" skeletonHeight="2" />
      </StyledFlex>
    );
  }

  if (isError) {
    return (
      <ErrorBanner
        description="Não foi possível obter o status do pedido"
        retry={refetch}
      />
    );
  }

  const status = Order.StatusData[data!.value];

  return (
    <a href={"/orders/" + data?.order_id}>
      <StyledFlex
        color={status.color}
        tabIndex={1}
        direction="column"
        justify="space-between"
      >
        <Flex alignItems="center" justify="space-between">
          <Flex alignItems="center">
            <Icon name={status.icon} />
            <Text marginInlineStart="10px" fontSize={"sm"}>
              {status.name}
            </Text>
          </Flex>
          <Flex direction="column" alignItems="end">
            <Text fontSize="2xs" justifySelf="flex-end">
              {data?.occurred_at.toLocaleString("pt-BR", {
                month: "numeric",
                day: "numeric",
                hour: "numeric",
                minute: "numeric",
              })}
            </Text>
            <Text fontSize="2xs" justifySelf="flex-end">
              há {Order.minuteDifference(data!.occurred_at)} minutos
            </Text>
          </Flex>
        </Flex>

        <Flex marginTop="8px" alignItems="center" justify="space-between">
          <Text fontSize={"xs"}>{data?.customer_name + " • " + phone}</Text>
          <ArrowForwardIcon />
        </Flex>
      </StyledFlex>
    </a>
  );
}

const StyledFlex = styled(Flex)`
  min-height: 5rem;
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

  a {
    font-size: 18px;
  }

  h4 {
    font-size: 24px;
  }
`;
