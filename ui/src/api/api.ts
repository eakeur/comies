export namespace API {
    export const idHeader = "Location"

    export type ErrorDesc = {
        code: string;
        message: string;
        target?: any;
    };

    export function authorize(init: RequestInit = {}): RequestInit {
        return init;
    }

    export function response(res: globalThis.Response) {
        if (!res.ok) {
            res.json().then((data) => {
                throw data;
            });
        }

        return res
    }

    type RouteOptions = { params?: { [key: string]: string }, query?: any }
    export function route(name: string, { params: prm, query: q }: RouteOptions = {}) {
        const route = process.env[`REACT_APP_${name}`]
        if (!route) {
            return ""
        }

        return query(params(route, prm), q)
    }

    function params(url: string, prm?: { [key: number]: string }){
        if (prm) Object.entries(prm).forEach(([k, v]) => url = url.replace(k, v))
        return url
    }
    function query(url: string, query?: any) {
        return !query ? url : `${url}?${new URLSearchParams(query).toString()}`
    }
}
