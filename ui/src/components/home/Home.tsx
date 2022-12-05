import styled from "styled-components";
import { OrdersOverview } from "./OrdersOverview";
import { HomeShortcuts } from "./Shortcuts";

const StyledMain = styled.main`
    padding: 20px;
    display: grid;
    grid-template-columns: 1.2fr 2fr 3fr;
    grid-gap: 50px;

    > div:nth-child(1) {
        position: sticky;
        top:0;
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
`

export function Home(){
    return (
        <StyledMain>
            <HomeShortcuts/>
            <OrdersOverview/>
        </StyledMain>
    )
}