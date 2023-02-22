import { Drawer, DrawerOverlay, DrawerContent, DrawerCloseButton, DrawerHeader, DrawerBody, DrawerFooter } from "@chakra-ui/react";
import { Context, useDrawerProvider } from "../../hooks/drawer";

export function DrawerProvider({children}: {children: React.ReactNode}){
    const [drawer, setDrawer] = useDrawerProvider()

    function onClose() {
        drawer?.options?.onClose?.call(undefined)
        setDrawer(undefined)
    }

    return (
        <Context.Provider value={setDrawer}>
            {children}
            <Drawer 
                size={drawer?.options?.size}
                isOpen={drawer?.body !== undefined ?? false} 
                placement={window.innerWidth < 820 ? "bottom" : drawer?.options?.placement as any} 
                onClose={onClose} 
                finalFocusRef={drawer?.options?.finalFocusRef}>
                <DrawerOverlay />
                <DrawerContent minWidth="40vw">
                    <DrawerCloseButton />
                    {
                        drawer?.header && <DrawerHeader>{drawer?.header}</DrawerHeader>
                    }
                    {
                        drawer?.body && <DrawerBody>{drawer?.body}</DrawerBody>
                    }
                    {
                        drawer?.footer && <DrawerFooter>{drawer?.footer}</DrawerFooter>
                    }
                </DrawerContent>
            </Drawer>
        </Context.Provider>
    )
}