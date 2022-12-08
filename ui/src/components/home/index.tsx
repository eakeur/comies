import { ArrowForwardIcon } from "@chakra-ui/icons";
import { Button, Flex, Input, Text } from "@chakra-ui/react";
import styled from "styled-components";
import { runAction } from "../../handlers/handler";
import { OrderFastChecker } from "./order-fast-checker";
import { OrderStatuses } from "./order-statuses";
import { HomeShortcutButton } from "./shortcut-button";

export function Home() {
  return (
    <StyledMain>
      <Flex wrap="wrap" direction="column">
        <Text as={"h1"} fontSize={"4xl"} mb="10px" fontWeight="bold">
          Comies
        </Text>

        <Input type={"text"} variant="filled" placeholder="Busque tudo..." />

        <Text as={"h4"} fontSize={"md"} mt="30px">
          O que vamos fazer hoje?
        </Text>
        <HomeShortcutButton
          name="Novo pedido"
          shortcut="Shift + A"
          icon="BuildQueueNew"
          color="#ef6c00"
          onClick={() => runAction("add_order")}
        />
        <HomeShortcutButton
          name="Novo produto"
          shortcut="Shift + D"
          icon="AppIconDefaultAdd"
          color="#1565c0"
          onClick={() => runAction("add_product")}
        />
        <HomeShortcutButton
          name="Verificar estoque"
          shortcut="Shift + V"
          icon="ProductionFloorManagement"
          color="#c62828"
          onClick={() => runAction("verify_stock")}
        />
        <HomeShortcutButton
          name="Nova entrada/saída"
          shortcut="Shift + M"
          icon="SwitcherStartEnd"
          color="#6a1b9a"
          onClick={() => runAction("add_stock_movement")}
        />
      </Flex>
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

        <Text as={"h3"} fontSize={"l"}>
          Verifique o status de um pedido:
        </Text>
        <OrderFastChecker></OrderFastChecker>

        <Text as={"h3"} fontSize={"l"}>
          Resumo do dia
        </Text>
        <OrderStatuses />
      </section>
    </StyledMain>
  );
}

const StyledMain = styled.main`
  padding: 20px;
  display: grid;
  grid-template-columns: 1.2fr 2fr 3fr;
  grid-gap: 50px;

  > div:nth-child(1) {
    position: sticky;
    top: 0;
  }

  > section:nth-last-child(1) {
    > * {
      margin-bottom: 30px;
    }

    h3 {
      margin-bottom: 5px;
    }
  }

  @media (max-width: 980px) {
    display: block;

    > * {
      margin-bottom: 20px;
    }

    > div:nth-child(1) {
      position: relative;
      flex-direction: row;
    }
  }
`;
