import { AddIcon } from "@chakra-ui/icons";
import { FormControl, FormLabel, Input, Text, FormHelperText, Stack, Flex, Button, ButtonGroup, IconButton, Select } from "@chakra-ui/react";
import styled from "styled-components";
import { useProductForm } from "../../hooks/product";
import { Product, ProductType, Unit } from "../../model/product";

const StyledForm = styled.form`
    & > div {
        margin-bottom: 30px;
    }
`

export function ProductForm(){

    const {control, submit} = useProductForm()

    return (
        <StyledForm onSubmit={submit}>
            <Stack id="naming-section" direction={["column", null, "row"]}>
                <FormControl id="type" w={["100%", null, "100%"]}>
                    <FormLabel fontSize="sm">Tipo</FormLabel>
                    <Select {...control.register("type", {valueAsNumber: true, required: true})}>
                        <option value={ProductType.none}>Nenhum</option>
                        <option value={ProductType.input}>Entrada</option>
                        <option value={ProductType.input_composite}>Entrada composta</option>
                        <option value={ProductType.output}>Saída</option>
                        <option value={ProductType.output_composite}>Saída composta</option>
                    </Select>
                </FormControl>
            </Stack>

            <Text>Identificação</Text>
            <Stack id="naming-section" direction={["column", null, "row"]}>
                <FormControl id="code" w={["100%", null, "30%"]}>
                    <FormLabel fontSize="sm">Código</FormLabel>
                    <Input {...control.register("code", {required: true, minLength: 3, maxLength: 6})}/>
                    <FormHelperText></FormHelperText>
                </FormControl>
                <FormControl id="name" w={["100%", null, "70%"]}>
                    <FormLabel fontSize="sm">Nome</FormLabel>
                    <Input {...control.register("name", {required: true, minLength: {value: 3, message: "sth"}, maxLength: 60})} placeholder=""/>
                </FormControl>
            </Stack>

            <Text>Vendas</Text>
            <Stack id="naming-section" direction={["column", null, "row"]}>
                <FormControl id="cost_price" w={["100%", null, "33%"]}>
                    <FormLabel fontSize="sm">Custo</FormLabel>
                    <Input {...control.register("cost_price", {valueAsNumber: true})} step="0.50" type="number"/>
                </FormControl>
                <FormControl id="sale_price" w={["100%", null, "33%"]}>
                    <FormLabel fontSize="sm">Preço</FormLabel>
                    <Input {...control.register("sale_price", {valueAsNumber: true})} step="0.50" type="number"/>
                </FormControl>
                <FormControl id="sale_unit" w={["100%", null, "33%"]}>
                    <FormLabel fontSize="sm">Unidade</FormLabel>
                    <Select {...control.register("sale_unit")}>
                        {
                            Object.entries(Unit).map(x => <option key={x[0]} value={x[1]}>{x[1]}</option>)
                        }
                    </Select>
                </FormControl>
            </Stack>

            <Text>Estoque</Text>
                    <Stack id="naming-section" direction={["column", null, "row"]}>
                        <FormControl id="minimum_quantity" w={["100%", null, "33%"]}>
                            <FormLabel fontSize="sm">Min. estoque</FormLabel>
                            <Input {...control.register("minimum_quantity", {valueAsNumber: true, deps: "type"})} step="0.50" type="number" readOnly={Product.isProductComposite(control.getValues().type)}/>
                        </FormControl>
                        <FormControl id="maximum_quantity" w={["100%", null, "33%"]}>
                            <FormLabel fontSize="sm">Max. estoque</FormLabel>
                            <Input {...control.register("maximum_quantity", {valueAsNumber: true, deps: "type"})} type="number" step="0.50" readOnly={Product.isProductComposite(control.getValues().type)}/>
                        </FormControl>
                        <FormControl id="location" w={["100%", null, "33%"]}>
                            <FormLabel fontSize="sm">Local</FormLabel>
                            <Input {...control.register("location")}/>
                        </FormControl>
                    </Stack>

            <Button width="100%" colorScheme='green' type="submit">
                Salvar
            </Button>
            
        </StyledForm>
    )
}