import {
  Box,
  Button,
  Text,
} from "@chakra-ui/react";
import { useForm } from "react-hook-form";
import styled from "styled-components";
import { CustomerForm } from "components/orders/place/customer-form";
import { ItemSearch } from "./item-search-form";
import { useState } from "react";
import { Order } from "core/order";
import { ItemForm } from "./item-form";
import { AnimatePresence } from "framer-motion";

const StyledForm = styled.form`
  & > div {
    margin-bottom: 30px;
  }
`;

export function PlaceOrder() {
  const [items, setItems] = useState<Order.Item[]>([])

  const form = useForm();

  return (
    <StyledMain>
      <StyledForm>

        <CustomerForm form={form} />

        <Box>
          <Text fontSize="x-large">Itens</Text>
          <div>
            <AnimatePresence>
            {
              items.map((it, i) => {
                return (
                  <ItemForm
                    key={i}
                    item={it}
                    onRemove={() => setItems(items.filter((_, idx) => i !== idx))} />
                )
              })
            }
            </AnimatePresence>
          </div>


          <ItemSearch onSelect={(it) => {
            const existing = items.findIndex(item => it.product.id === item.product.id)

            if (existing === -1)
              return setItems(items.concat(it))

            const actual = items[existing]
            setItems([
              ...items.slice(0, existing),
              {
                ...actual,
                quantity: actual.quantity + it.quantity
              },
              ...items.slice(existing + 1),
            ])
          }} />
        </Box>

        <Button width="100%" colorScheme="green" type="submit">
          Salvar
        </Button>
      </StyledForm>
    </StyledMain>
  );
}



const StyledMain = styled.main`
  padding: 20px;
`;