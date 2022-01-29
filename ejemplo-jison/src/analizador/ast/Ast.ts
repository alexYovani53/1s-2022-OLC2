import { Instruccion } from "./abstract/Instruccion";


export class Ast {


    public instrucciones: Array<Instruccion>;

    constructor(instrucciones: Array<Instruccion>){
        this.instrucciones = instrucciones;
    }


}