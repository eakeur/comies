import { Text } from "@chakra-ui/react";

export function Price({ value, currency = "R$" }: { value: number, currency?: string }) {
    return (
        <Text as="span" fontSize="lg" textAlign="left" display="flex" alignItems="center" minW="100px">
            <Text as="span" fontSize="sm" marginRight="10px">{currency}</Text>
            {(value / 100).toFixed(2)}
        </Text>
    )
}

export function Quantity({ value, unit = "un" }: { value: number, unit?: string }) {
    return (
        <Text as="span" fontSize="lg" textAlign="left" display="flex" alignItems="center" minW="100px">
            <Text as="span" fontSize="sm" marginRight="10px">{unit}</Text>
            {value}
        </Text>
    )
}