import { Flex, Input, Text } from "@chakra-ui/react";
import styled from "styled-components";
import { runAction } from "handlers/handler";
import { MyOrders } from "./my-orders";
import { HomeShortcutButton } from "./shortcut-button";
import { OrderForm } from "components/product/order-form";

export function Home() {

  return (
    <StyledMain>
      <Flex wrap="wrap" direction="column">
        <Text as={"h1"} fontSize={"4xl"} mb="10px" fontWeight="bold">
          Comies
        </Text>

        <Input type={"text"} variant="filled" placeholder="Busque tudo..." />

        <Text as={"h4"} fontSize={"md"} mt="30px" id="shortcuts-title">
          O que vamos fazer hoje?
        </Text>
        <Flex id="shortcuts" direction="column">
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

      </Flex>
      <MyOrders />
      <OrderForm/>
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

  section#myorders {
    > * {
      margin-bottom: 30px;
    }
  }

  div#shortcuts {
    @media (max-width: 980px) {
      background-color: #FEFBEA;
      position: fixed;
      bottom: 0;
      left: 0;
      right: 0;
      box-sizing: border-box;
      flex-direction:row;
    }
  }

  @media (max-width: 980px) {
    display: block;

    h4#shortcuts-title {
      display: none;
      
    }

    > * {
      margin-bottom: 20px;
    }

    > div:nth-child(1) {
      position: relative;
      flex-direction: row;
    }
  }
`;
