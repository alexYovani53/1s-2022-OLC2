import { NodoAst } from "./NodoAst";
import { Retorno } from "./Retorno";




export abstract class Expresion extends  NodoAst{

    constructor(line:number, columna:number){
        super(line,columna);
    }

    public abstract getValorImplicito():Retorno;

}