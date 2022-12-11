import { Flex, Text } from "@chakra-ui/react";
import { darken, lighten } from "polished";
import styled from "styled-components";
import { Icon } from "./Icon";

type ErrorBannerProps = {
  title?: string;
  description: string;
  retry?: () => void;
};
export function ErrorBanner(props: ErrorBannerProps) {
  return (
    <StyledOrderStatusError justify="space-between"
      tabIndex={1}
    >
      <Flex>
        <Flex direction="column"
          justify="space-between">
          <Text fontSize="sm" fontWeight="bold">{props.title ?? "Um erro aconteceu 😔"}</Text>
          <Text fontSize="sm">{props.description}</Text>
        </Flex>
      </Flex>
      {props.retry && (
        <Flex justifyContent="end" alignItems="center">
          <Text fontSize="xs" marginRight="10px">Tentar novamente</Text>
          <button onClick={props.retry}>
            <Icon name="Refresh" />
          </button>
        </Flex>
      )}

    </StyledOrderStatusError>
  );
}

const StyledOrderStatusError = styled(Flex)`
  padding: 20px;
  border-radius: 7px;
  border: 2px solid ${(_) => darken(0.2, "#ff0033")};
  background: ${(_) => lighten(0.45, "#ff0033")};

  * {
    color: ${(_) => darken(0.35, "#ff0033")};
    margin-block: 2px;
  }
`;
