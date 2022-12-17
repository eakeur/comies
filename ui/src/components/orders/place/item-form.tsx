import { DeleteIcon } from "@chakra-ui/icons";
import { Flex, FormControl, Input, Stack, IconButton, } from "@chakra-ui/react";
import { ListAddingAnimation } from "components/animations/animations";
import { Order } from "core/order";
import { motion } from "framer-motion";

import { useForm } from "react-hook-form";

interface Props {
    item: Order.Item
    onRemove: () => void,
}

export function ItemForm({ item, onRemove }: Props) {
    const { control, register } = useForm({ defaultValues: item });

    return (
        <MotionStack
            id={item.product.id}
            direction="row"
            paddingBottom="20px"
            initial="hidden"
            animate="visible"
            variants={ListAddingAnimation.parent}
            exit={ListAddingAnimation.parent.hidden}>
            <FormControl width="40%">
                <Input disabled={true} value={item.product.name} />
            </FormControl>
            <FormControl width="10%">
                <Input disabled={true} value={item.product.price} />
            </FormControl>
            <FormControl width="10%">
                <Input autoFocus={item.quantity !== undefined} {...register("quantity", {
                    valueAsNumber: true,
                })} defaultValue={item.quantity.toString()} />
            </FormControl>
            <FormControl width="10%">
                <Input {...register("discounts")} />
            </FormControl>
            <Flex>
                <IconButton
                    colorScheme="red"
                    aria-label={"Remover produto do pedido"}
                    icon={<DeleteIcon />}
                    onClick={onRemove} />
            </Flex>
        </MotionStack>
    )
}

const MotionStack = motion(Stack)