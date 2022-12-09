import { lighten } from "polished";
import styled from "styled-components";
import { Icon } from "components/shared/Icon";

export function HomeShortcutButton({
  name,
  icon,
  onClick,
  color,
}: HomeShortcutButtonProps) {
  return (
    <StyledHomeShortcutButton onClick={onClick} color={color} title={name}>
      {icon && <Icon name={icon} />}
      <span>{name}</span>
    </StyledHomeShortcutButton>
  );
}

type HomeShortcutButtonProps = {
  icon?: string;
  color?: string;
  name: string;
  shortcut: string;
  onClick?: () => void;
};

const StyledHomeShortcutButton = styled.button`
  display: flex;
  flex-direction: row;
  align-items: center;

  width: 100%;
  height: 55px;

  padding: 10px;
  margin: 5px;
  border-radius: 5px;

  text-align: left;
  font-size: 0.9rem;
  font-weight: 500;
  color: var(--chakra-colors-orange-600);

  transition: all 200ms ease-out;

  i {
    font-size: 24px;
    text-align: center;
    margin-right: 15px;
  }

  &:hover {
    background: ${() => lighten(0.25, "#ffb74d")};
  }
  &:active {
    background: ${() => lighten(0.2, "#ffb74d")};
  }
`;