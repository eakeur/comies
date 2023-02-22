import { DeleteIcon, EditIcon } from "@chakra-ui/icons";
import {
    Box, FormControl, FormLabel, Input,
    Text,
    useDisclosure, Button, Modal, ModalBody, ModalContent, ModalFooter, ModalHeader, ModalOverlay, Divider, UseDisclosureReturn, Flex, IconButton, useMediaQuery

} from "@chakra-ui/react";
import { Price, Quantity } from "components/shared/price";
import { Ordering } from "core/order";
import { useState } from "react";
import { ItemQuantitySetter } from "./quantity";






export const TicketItemEditorModal = ({ item, disclosure, add }: { item: Ordering.TicketItem, disclosure: UseDisclosureReturn, add: (item: Ordering.TicketItem) => void }) => {
    const { isOpen, onClose } = disclosure

    const [obs, setObs] = useState(item.observations)

    const [qty, setQty] = useState(item.quantity)

    return (
        <Modal size="lg" isOpen={isOpen} onClose={onClose}>
            <ModalOverlay />
            <ModalContent>
                <ModalHeader>Detalhes do item</ModalHeader>
                <ModalBody>
                    <Box display="grid" gridTemplateColumns="0.7fr 0.3fr" justifyContent="space-between" gridGap="10px">
                        <SaleableItem saleable={item.saleable} showStock />
                        <Box>
                            <Text fontSize="xs">Total</Text>
                            <Price value={qty * item.saleable.price} />
                        </Box>
                    </Box>
                    <Divider marginBlock="10px" />
                    <FormControl id="quantity" marginBlock="10px">
                        <FormLabel fontSize="sm">Quantidade</FormLabel>
                        <ItemQuantitySetter quantity={qty} onChange={setQty} />
                    </FormControl>
                    <FormControl id="observations">
                        <FormLabel fontSize="sm">Observações</FormLabel>
                        <Input defaultValue={obs} onChange={(ev) => setObs(ev.target.value)} />
                    </FormControl>
                </ModalBody>

                <ModalFooter>
                    <Button colorScheme='blue' mr={3} onClick={() => {
                        add({
                            ...item,
                            quantity: qty,
                            observations: obs
                        })
                        onClose()
                    }}>Adicionar</Button>
                    <Button colorScheme='red' variant='outline' onClick={onClose}>Cancelar</Button>
                </ModalFooter>
            </ModalContent>
        </Modal>
    )

}

export const SaleableItem = ({ saleable, showStock = false }: { saleable: Ordering.SaleableItem, showStock?: boolean }) => (
    <Box display="grid" gridTemplateColumns="0.7fr 0.3fr" justifyContent="space-between" gridGap="10px">
        <Box>
            <Text><b>{saleable.code}</b> - {saleable.name}</Text>
            {
                showStock && <Text fontSize="sm">Estoque disponível: {saleable.stock}</Text>
            }
        </Box>
        <Price value={saleable.price} />
    </Box>
)

export const TicketItem = ({ item, onRemove, onSave }: {
    item: Ordering.TicketItem
    onRemove?(): void,
    onSave?(it: Ordering.TicketItem): void
}) => {

    const [isPhone] = useMediaQuery('(max-width: 800px)')

    const disclosure = useDisclosure()

    if (isPhone) {
        <Flex onClick={disclosure.onOpen}>
            <Box>
                <Text>{item.saleable.name}</Text>
                <Text>{item.observations}</Text>
            </Box>
            <Flex>
                <Box>
                    <Text fontSize="xs">Quantidade</Text>
                    <Quantity value={item.quantity} />
                </Box>
                <Box>
                    <Text fontSize="xs">Total</Text>
                    <Price value={item.quantity * item.saleable.price} />
                </Box>
            </Flex>
            <Box>
                <Text fontSize="xs">Detalhes</Text>
                <Text as="span" fontSize="sm" marginRight="10px">{item.observations}</Text>
            </Box>
            <Box>
                <Text fontSize="xs">Unidade</Text>
                <Price value={item.saleable.price} />

            </Box>
            <Box>
                <Text fontSize="xs">Total</Text>
                <Price value={item.quantity * item.saleable.price} />
            </Box>
        </Flex>
    }

    return (
        <Box display="grid" gridTemplateColumns="3fr 1fr 1fr 2fr 2fr 2fr" justifyContent="space-between" gridGap="10px" alignItems="center">
            <Text><b>{item.saleable.code}</b><br />{item.saleable.name}</Text>

            <Box>
                <Text fontSize="xs">Detalhes</Text>
                <Text as="span" fontSize="sm" marginRight="10px">{item.observations}</Text>
            </Box>
            
            <Box>
                <Text fontSize="xs">Quantidade</Text>
                <Quantity value={item.quantity} />
            </Box>
            <Box>
                <Text fontSize="xs">Unidade</Text>
                <Price value={item.saleable.price} />

            </Box>
            <Box>
                <Text fontSize="xs">Total</Text>
                <Price value={item.quantity * item.saleable.price} />
            </Box>
            {
                onSave && onRemove &&
                <Flex justifySelf="end">
                    <IconButton
                        colorScheme="blue"
                        aria-label={"Remover produto do pedido"}
                        icon={<EditIcon />}
                        onClick={disclosure.onOpen} />
                    <IconButton
                        marginLeft="5px"
                        colorScheme="red"
                        aria-label={"Remover produto do pedido"}
                        icon={<DeleteIcon />}
                        onClick={onRemove} />
                    <TicketItemEditorModal item={item} disclosure={disclosure} add={onSave} />
                </Flex>
            }
        </Box>
    )
}