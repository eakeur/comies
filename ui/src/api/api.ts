import { fail } from "core/failures";

export namespace API {
    export const idHeader = "Location"
    
    export function authorize(init: RequestInit = {}): RequestInit {
        return init;
    }

    export function handle(res: globalThis.Response) {
        if (!res.ok) {
            res.json().then((data) => {
                throw data;
            }).catch(fail);
        }

        return res
    }

    export function json(res: globalThis.Response) {
        return res.json().catch(fail);
    }

    class URL extends String {
        params(key: string, value: string): URL {
            return new URL(this.replace(key, value));
        }

        query(query: any): URL {
            return !query ? this : new URL(`${this}?${new URLSearchParams(query).toString()}`)
        }

        request(init?: RequestInit){
            return fetch(this.toString(), authorize(init)).then(handle)
        }
    }

    export function route(name: string) {
        return new URL(name);
    }

    
}
