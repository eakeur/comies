class Failure extends Error {
    readonly code: string;
    readonly title: string;
    readonly description?: string;

    constructor(code: string, title: string, description?: string, err?: any) {
        super(code);

        this.code = code;
        this.title = title;
        this.description = description;
        this.cause = err;
    }

    is(code: string) {
        return code === this.code;
    }

    withCause(cause: any) {
        return new Failure(this.code, this.title, this.description, cause)
    }
}

const errors = [
    new Failure(
        "ERR_INTERNAL_SERVER_ERROR",
        "Eita, um erro aconteceu",
        "Tente novamente ou contate um administrador",
    ),
    new Failure(
        "PRODUCT_NOT_FOUND",
        "Eita! Não encontramos esse produto",
        "Ele não existe ou não pôde ser encontrado com as informações passadas."
    ),
    new Failure(
        "PRODUCT_CODE_ALREADY_EXISTS",
        "Ops! Já existe um produto com este código",
        "Tente criá-lo com outro código."
    ),
] as const

export function fail(err?: any) {
    if (err.code) throw (errors.find(x => x.is(err.code)) ?? errors[0]).withCause(err) 

    console.error(err)
    throw errors[0].withCause(err)
}

export function as(err: any){
    if (err instanceof Failure){
        return err
    }

    return errors[0].withCause(err)
}