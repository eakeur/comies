import styled from "styled-components"

const StyledInput = styled.input`
`

export function Input(props: any){
    return (
        <StyledInput {...props}/>
    )
}