import { Expresion } from "../abstract/Expresion";
import { Retorno, tipo } from "../abstract/Retorno";


export class Primitivo extends Expresion{


    private valor:any;
    private tipo:tipo;


    constructor(valor:any, tipo:tipo, linea:number, columna:number){
        super(linea,columna);
        this.valor = valor;
        this.tipo = tipo;
    }
    
    public getValorImplicito(): Retorno {
        return {value:this.valor , type:this.tipo};
    }
    
}