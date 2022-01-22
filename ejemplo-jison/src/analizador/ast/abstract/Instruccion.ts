import { NodoAst } from "./NodoAst";


export abstract class Instruccion extends NodoAst{

    constructor(line:number, columna:number){
        super(line,columna);
    }


    public abstract ejecutar():any;

}