type Action = {name: string, handler: () => void}

let actions: Action[] = []

export function addAction(action: Action) {
    if (action.name === "") {
        return
    }

    actions = [...actions, action];
}

export function runAction(actionName: string) {
    if (actionName === "") {
        return
    }

    actions.find(a => a.name === actionName)?.handler()
}