import { createContext, useContext, useState } from "react"

type Drawer = {
    header?: JSX.Element,
    body?: JSX.Element,
    footer?: JSX.Element,
    options?: {
        placement?: string 
        onClose?: () => void,
        finalFocusRef?: React.RefObject<any>
    },
} | undefined

export const Context = createContext<React.Dispatch<React.SetStateAction<Drawer>>>((v) => v);

export function useDrawer(){
    return useContext(Context);
}

export function useDrawerProvider(){
    return useState<Drawer>({options: {}})
}