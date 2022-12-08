import { Button, Input, InputGroup, InputRightElement } from "@chakra-ui/react";
import { useState } from "react";
import styled from "styled-components";
import { OrderBanner } from "./order-banner";

export function OrderFastChecker() {
  const [phone, setPhone] = useState("");

  return (
    <StyledForm
      onSubmit={(ev) => {
        const phone = new FormData(ev.currentTarget).get("phone");
        if (phone) setPhone(phone.toString());

        ev.preventDefault();
      }}
    >
      <InputGroup size="md">
        <Input
          pr="4.5rem"
          type="phone"
          name="phone"
          placeholder="Telefone do cliente"
          variant="outline"
          defaultValue={phone}
        />
        <InputRightElement width="7rem">
          <Button
            type="submit"
            h="1.75rem"
            size="sm"
            aria-label="Pesquisar pedido por telefone do cliente"
          >
            Pesquisar
          </Button>
        </InputRightElement>
      </InputGroup>
      {phone !== "" && <OrderBanner phone={phone} />}
    </StyledForm>
  );
}

const StyledForm = styled.form`
  > a,
  > div:nth-last-child(1) {
    display: block;
    margin-top: 20px;
  }
`;
