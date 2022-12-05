import { useCallback, useEffect } from "react";
import { runAction } from "../handlers/handler";

const KeyActions = {
    "A": "add_order",
    "D": "add_product",
    "V": "verify_stock",
    "M": "add_stock_movement",
}

export function useShortcuts() {
    const handler = useCallback((event: KeyboardEvent) => {
        if (!event.shiftKey && event.key === "") {
            return
        }

        const action = KeyActions[(event.key as keyof typeof KeyActions)]
        if (!action) {
            return
        }

        runAction(action)
    }, []);
    
    useEffect(() => {
        document.addEventListener('keydown', handler);
        return () => document.removeEventListener('keydown', handler);
    }, [handler]);

    return {...KeyActions}
}