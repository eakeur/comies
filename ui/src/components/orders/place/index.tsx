import {
  Button,
  Drawer,
  DrawerBody,
  DrawerHeader,
  DrawerOverlay,
  DrawerContent,
  DrawerCloseButton,
  useDisclosure,
  Box,
  FormControl,
  FormLabel,
  Input,
  UnorderedList,
  Text,
  Flex,
  FormHelperText,
  Stack,
  Textarea,
  RadioGroup,
  Radio,
} from "@chakra-ui/react";
import { useRef } from "react";
import { useForm } from "react-hook-form";
import { useSaleableItemsSearch, useTicketItems } from "./hook";
import { TicketItem } from "./saleable-item";
import { SearchResult } from "./search-result";

export function PlaceOrder() {
  const { isOpen, onOpen, onClose } = useDisclosure()

  const { data, search, setSearch } = useSaleableItemsSearch()

  const items = useTicketItems()

  const { register } = useForm();

  const nextRef = useRef(null);

  return (
    <form>
      <Box marginBottom="20px">
        <Text fontSize="x-large">Cliente</Text>
        <Stack id="customer-data" direction="column">
          <FormControl id="customer_name">
            <FormLabel fontSize="sm">Nome</FormLabel>
            <Input {...register("customer_name")} autoFocus={true} />
            <FormHelperText></FormHelperText>
          </FormControl>
          <FormControl id="customer_phone">
            <FormLabel fontSize="sm">Telefone</FormLabel>
            <Input {...register("customer_phone")} />
            <FormHelperText>Necessário em caso de entrega</FormHelperText>
          </FormControl>
          <FormControl id="customer_address">
            <FormLabel fontSize="sm">Endereço</FormLabel>
            <Input {...register("customer_address")} />
            <FormHelperText>Necessário em caso de entrega</FormHelperText>
          </FormControl>
        </Stack>
      </Box>

      {
        items.list.length > 0 ?
          <Box overflowY="auto" marginBottom="20px">
            {
              items.list.map((it, i) => {
                return (
                  <TicketItem key={i} item={it} onSave={(it) => items.setItem(it, i)} onRemove={() => items.removeItem(i)} />
                )
              })
            }
          </Box> :
          <Flex alignSelf="center" justifyContent="center" marginBottom="20px">
            <Text>Adicione pelo menos um item ao pedido</Text>
          </Flex>
      }

      <Box marginBottom="20px">
        <Button onFocus={onOpen} onClick={onOpen} w="100%">Adicionar itens</Button>
        <Drawer
          isOpen={isOpen}
          placement='left'
          size="md"
          onClose={onClose}
          finalFocusRef={nextRef}
        >
          <DrawerOverlay />
          <DrawerContent>
            <DrawerCloseButton onClick={onClose} />
            <DrawerHeader>
              <FormControl paddingBlockEnd={5}>
                <FormLabel fontSize="sm">Pesquisar items</FormLabel>
                <Input autoFocus={true} defaultValue={search} onChange={(ev) => setSearch(ev.target.value)} />
              </FormControl>
            </DrawerHeader>

            <DrawerBody>
              <UnorderedList>
                {
                  data?.map((prod, i) => <SearchResult
                    key={i}
                    saleable={prod}
                    onAdd={items.addItem} />)
                }
              </UnorderedList>
            </DrawerBody>
          </DrawerContent>
        </Drawer>
      </Box>

      <Box marginBottom="20px">
        <Text fontSize="x-large">Detalhes</Text>
        <Stack id="order-detail" direction="column">
          <FormControl id="observations">
            <FormLabel fontSize="sm">Observações</FormLabel>
            <Textarea {...register("customer_address")} ref={nextRef} />
          </FormControl>
          <RadioGroup defaultValue='2'>
            <FormLabel fontSize="sm">Forma de entrega</FormLabel>
            <Stack spacing={5} direction='row'>
              <Radio colorScheme='red' value='1'>
                Retirada
              </Radio>
              <Radio colorScheme='green' value='2'>
                Entrega
              </Radio>
            </Stack>
          </RadioGroup>
        </Stack>
      </Box>

      <Button width="100%" colorScheme="green" type="submit">
        Confirmar pedido
      </Button>
    </form>
  );
}