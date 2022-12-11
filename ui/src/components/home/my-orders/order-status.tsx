import { Flex, SkeletonText, Text, useToast } from "@chakra-ui/react";
import { getOrderCountByStatus } from "api/ordering";
import { Icon } from "components/shared/Icon";
import { Order } from "core/order";
import { useQuery } from "react-query";
import { darken } from "polished";
import styled from "styled-components";
import { motion } from "framer-motion";
import { ListShowUpAnimation } from "components/animations/animations";

export function OrderCountByStatus({ status }: { status: number }) {

  const toast = useToast()
  const {
    isLoading,
    isError,
    data,
    refetch,
  } = useQuery(["statuscount", status], () =>
    getOrderCountByStatus(status)
  , {
    onError(err){
      if (!toast.isActive("statuscount")){
        toast({
          title: "Não foi possível recuperar os pedidos",
          status: "error",
          id: "statuscount"
        })
      }
    }
  });

  const { icon, color, name } = Order.StatusData[status];

  return (
    <StatusCountBanner
      color={color}
      tabIndex={1}
      direction="column"
      justify="space-between"
      variants={ListShowUpAnimation.children}
    >
      <Icon name={icon} />
      {(function () {
        if (isLoading) {
          return <SkeletonText
          startColor={darken(0.35, color)}
            endColor={darken(0.05, color)}
            mt="4"
            noOfLines={2}
            spacing="2"
            skeletonHeight="2"
          />;
        }

        if (isError) {
          return <Flex>
            <Text fontSize="xs" marginRight="10px">Tentar novamente</Text>
            <button onClick={() => refetch()}>
              <Icon name="Refresh" />
            </button>
          </Flex>
        }

        return <>
          <h4>{data}</h4>
          <Text fontSize="sm">{name}</Text>
        </>;
      })()}
      
    </StatusCountBanner>
  );
}

const StatusCountBanner = styled(motion(Flex))`
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
