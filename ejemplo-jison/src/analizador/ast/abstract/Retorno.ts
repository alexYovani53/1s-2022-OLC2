
export enum tipo{
    INTEGER = 0,
    FLOAT = 1,
    STRING = 2,
    BOOLEAN = 3,
    NULL= 4
}

export type Retorno={
    value:any,
    type:tipo
}