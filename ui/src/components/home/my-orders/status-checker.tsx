import { useQuery } from "react-query";
import { Box, Button, Flex, Input, InputGroup, InputLeftElement, InputRightElement, SkeletonCircle, SkeletonText, Text } from "@chakra-ui/react";
import { darken } from "polished";
import styled from "styled-components";
import { Ordering } from "core/order";
import { getOrderStatusByCustomerPhone } from "api/ordering";
import { Icon } from "../../shared/Icon";
import { ArrowForwardIcon } from "@chakra-ui/icons";
import { useEffect, useState } from "react";
import { OrderStatusRenderingData } from "components/shared/rendering";

export function OrderStatusChecker() {

  const [phone, setPhone] = useState("");

  const { data, isError, isLoading, refetch } = useQuery(phone, () => {
    return getOrderStatusByCustomerPhone(phone)
  }, { enabled: false });

  useEffect(() => {
    if (phone.length >= 9) {
      refetch()
    }
  }, [phone, refetch])


  const status = OrderStatusRenderingData[data?.value ?? Ordering.OrderStatus.pending];

  if (isError) {
    return (
      <StyledFlex color="#ececec">
        err
      </StyledFlex>
    );
  }

  if (isLoading) {
    return (
      <StyledFlex
        color={status.color}
        tabIndex={1}
        direction="column"
        justify="space-between"
      >
        <SkeletonCircle size="7" />
        <SkeletonText mt="4" noOfLines={1} spacing="4" skeletonHeight="2" />
      </StyledFlex>
    );
  }

  return (
    <StyledFlex
      color={status.color}
      tabIndex={1}
      direction="column"
      justify="space-between"
    >
      <form onSubmit={(ev) => {
        ev.preventDefault()

        const phone = new FormData(ev.currentTarget).get("phone")?.toString() ?? ""
        setPhone(phone)
      }}>
        <InputGroup size="md">
          <InputLeftElement>
            <Icon name="Phone"/>
          </InputLeftElement>
          <Input
            pr="4.5rem"
            type="phone"
            name="phone"
            placeholder="Telefone do cliente"
            borderBottomColor={darken(0.35, status.color)}
            variant="flushed"
            defaultValue={phone}
          />
          <InputRightElement width="7rem">
            <Button
              type="submit"
              h="1.75rem"
              size="sm"
              backgroundColor={status.color}
              variant="solid"
              aria-label="Pesquisar pedido por telefone do cliente"
            >
              Pesquisar
            </Button>
          </InputRightElement>
        </InputGroup>
      </form>
      {
        data && !isError && !isLoading && phone.length >= 9 && 
        <Box marginTop="15px">
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
                há {Ordering.minuteDifference(data!.occurred_at)} minutos
              </Text>
            </Flex>
          </Flex>

          <Flex marginTop="8px" alignItems="center" justify="space-between">
            <Text fontSize={"xs"}>{data?.customer_name}</Text>
            <ArrowForwardIcon />
          </Flex>
        </Box>
      }
    </StyledFlex>
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
