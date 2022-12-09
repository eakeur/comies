import { Flex, SkeletonText, Text } from "@chakra-ui/react";
import { getOrderCountByStatus } from "api/ordering";
import { Icon } from "components/shared/Icon";
import { Order } from "core/order";
import { useQuery } from "react-query";
import { darken } from "polished";
import styled from "styled-components";

export function OrderCountByStatus({ status }: { status: number }) {
  const {
    isLoading,
    data,
  } = useQuery(["statuscount", status], () =>
    getOrderCountByStatus(status)
  );

  const { icon, color, name } = Order.StatusData[status];

  return (
    <StatusCountBanner
      color={color}
      tabIndex={1}
      direction="column"
      justify="space-between"
    >
      <Icon name={icon} />
      {(function () {
        if (isLoading) {
          <SkeletonText
            endColor="white.200"
            mt="4"
            noOfLines={2}
            spacing="1"
            skeletonHeight="1"
          />;
        }

        return <h4>{data}</h4>;
      })()}
      <Text fontSize="sm">{name}</Text>
    </StatusCountBanner>
  );
}

const StatusCountBanner = styled(Flex)`
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
