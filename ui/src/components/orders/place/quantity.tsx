import { Button, HStack, Input, useNumberInput } from "@chakra-ui/react"

type Props = {
    quantity: number,
    onChange: (val: number) => void
}

export function ItemQuantitySetter({ quantity, onChange }: Props) {

    const { getInputProps, getIncrementButtonProps, getDecrementButtonProps } =
        useNumberInput({
            step: 1,
            defaultValue: quantity,
            min: 1,
            onChange: (s, n) => onChange(n)
        })

    const inc = getIncrementButtonProps()
    const dec = getDecrementButtonProps()
    const input = getInputProps()

    return (
        <HStack>
            <Button {...inc}>+</Button>
            <Input {...input} autoFocus={true} placeholder="Quantidade" />
            <Button {...dec}>-</Button>
        </HStack>
    )
}