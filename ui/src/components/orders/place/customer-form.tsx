import {
    FormControl,
    FormLabel,
    Input,
    Text,
    FormHelperText,
    Stack,
    Box,
} from "@chakra-ui/react";
import { UseFormReturn } from "react-hook-form";

interface Props {
    form: UseFormReturn
}

export function CustomerForm({form: {register}}: Props) {

    return (
        <Box>
            <Text fontSize="x-large">Cliente</Text>
            <Stack id="customer-data" direction="column">
                <FormControl id="customer_name">
                    <FormLabel fontSize="sm">Nome</FormLabel>
                    <Input {...register("customer_name")}/>
                    <FormHelperText></FormHelperText>
                </FormControl>
                <FormControl id="customer_phone">
                    <FormLabel fontSize="sm">Telefone</FormLabel>
                    <Input {...register("customer_phone")}/>
                    <FormHelperText>Necessário em caso de entrega</FormHelperText>
                </FormControl>
                <FormControl id="customer_address">
                    <FormLabel fontSize="sm">Endereço</FormLabel>
                    <Input {...register("customer_address")}/>
                    <FormHelperText>Necessário em caso de entrega</FormHelperText>
                </FormControl>
            </Stack>
        </Box>
    );
}
