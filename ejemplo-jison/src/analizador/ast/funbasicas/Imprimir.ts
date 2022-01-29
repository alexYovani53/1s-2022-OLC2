import { Expresion } from "../abstract/Expresion";
import { Instruccion } from "../abstract/Instruccion";
import { Retorno } from "../abstract/Retorno";



export class Imprimir extends Instruccion
{

    /*
    System.out.prinln("cadena1","cadena2","cadena3")
    
    */


    expresion:Expresion;

    constructor( expresion: Expresion , linea:number, columna:number){
        super(linea,columna)
        this.expresion = expresion;
    }

    public ejecutar(): any {


        let val: Retorno =  this.expresion.getValorImplicito();

        return val.value + "\n";
    }


    
}