import {
  FormControl,
  FormLabel,
  Input,
  Text,
  FormHelperText,
  Stack,
  Button,
  Menu,
  MenuButton,
} from "@chakra-ui/react";
import styled from "styled-components";
import { useProductForm } from "../../hooks/product";
import { OrderItemSearcher } from "./order-item-searcher";

const StyledForm = styled.form`
  & > div {
    margin-bottom: 30px;
  }
`;

export function OrderForm() {
  const { control, submit } = useProductForm();

  return (
    <StyledForm onSubmit={submit}>

      <Text fontSize="x-large">Cliente</Text>
      <Stack id="customer-data" direction="column">
        <FormControl id="customer_name">
          <FormLabel fontSize="sm">Nome</FormLabel>
          <Input/>
          <FormHelperText></FormHelperText>
        </FormControl>
        <FormControl id="customer_phone">
          <FormLabel fontSize="sm">Telefone</FormLabel>
          <Input/>
          <FormHelperText>Necessário em caso de entrega</FormHelperText>
        </FormControl>
        <FormControl id="customer_address">
          <FormLabel fontSize="sm">Endereço</FormLabel>
          <Input/>
          <FormHelperText>Necessário em caso de entrega</FormHelperText>
        </FormControl>
      </Stack>

      <OrderItemSearcher/>

      <Button width="100%" colorScheme="green" type="submit">
        Salvar
      </Button>
    </StyledForm>
  );
}
