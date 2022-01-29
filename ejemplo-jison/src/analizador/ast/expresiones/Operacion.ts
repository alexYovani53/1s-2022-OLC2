
import { Expresion } from "../abstract/Expresion";
import { Retorno, tipo } from "../abstract/Retorno";
import { Primitivo } from "./Primitivo";



export enum operador{
    MAS,
    MENOS, 
    MULTIPLICACION,
    DIVISION
}

/**
 *     operando  int    float    string   boolean    null
 * int          [      |       |         |         |     ]
 * float        [      |       |         |         |     ]
 * string       [      |       |         |         |     ]
 * boolean      [      |       |         |         |     ]
 * null         [      |       |         |         |     ]
 */

 const sumaDominante  = [
    [tipo.INTEGER, tipo.FLOAT, tipo.STRING, tipo.NULL,   tipo.NULL ],
    [tipo.FLOAT,   tipo.FLOAT, tipo.STRING, tipo.NULL,   tipo.NULL],
    [tipo.STRING,  tipo.STRING,tipo.STRING, tipo.STRING, tipo.NULL],
    [tipo.NULL  ,  tipo.NULL,  tipo.STRING, tipo.NULL,   tipo.NULL],
    [tipo.NULL,    tipo.NULL,  tipo.NULL,   tipo.NULL,   tipo.NULL],
]


const multi_division_Dominante:tipo [][] = [
    [tipo.INTEGER, tipo.FLOAT, tipo.NULL, tipo.NULL ,tipo.NULL],
    [tipo.FLOAT,   tipo.FLOAT, tipo.NULL, tipo.NULL,tipo.NULL],
    [tipo.NULL,    tipo.NULL,  tipo.NULL, tipo.NULL,tipo.NULL],
    [tipo.NULL,    tipo.NULL,  tipo.NULL, tipo.NULL,tipo.NULL],
    [tipo.NULL,    tipo.NULL,  tipo.NULL, tipo.NULL,tipo.NULL],
]

const restaDominante:tipo [][] = [
    [tipo.INTEGER,tipo.FLOAT, tipo.NULL, tipo.NULL ,tipo.NULL],
    [tipo.FLOAT,  tipo.FLOAT, tipo.NULL, tipo.NULL, tipo.NULL],
    [tipo.NULL,   tipo.NULL,  tipo.NULL, tipo.NULL, tipo.NULL],
    [tipo.NULL,   tipo.NULL,  tipo.NULL, tipo.NULL, tipo.NULL],
    [tipo.NULL,   tipo.NULL,  tipo.NULL, tipo.NULL, tipo.NULL],
]
export class Operacion extends Expresion{

    operandoIz: Expresion;
    operandoDe: Expresion;
    operador: operador;
    unario: boolean;

    constructor(opIz: Expresion, operador: operador, opDer: Expresion, unario:boolean, linea:number, columna:number){
        super(linea,columna);

        this.operandoIz  = opIz;
        this.operandoDe = opDer;
        this.operador = operador;
        this.unario = unario;
    }

  


    public getValorImplicito(): Retorno {

        let valorIz: Retorno;
        let valorDer: Retorno;
        let valorUn:Retorno;

        if(this.unario){
            valorUn  = this.operandoIz.getValorImplicito();
    
        }else{
            valorIz = this.operandoIz.getValorImplicito();
            valorDer = this.operandoDe.getValorImplicito();
        }


        let dominante:tipo;

        switch(this.operador){

            case operador.MAS:{

                
                dominante = sumaDominante[valorIz.type][valorDer.type];

                if(dominante == tipo.INTEGER) return {value:(valorIz.value + valorDer.value),type:tipo.INTEGER}
                else if(dominante == tipo.FLOAT ) return {value:(valorIz.value + valorDer.value),type:tipo.FLOAT};
                else if(dominante == tipo.STRING ) return {value:(valorIz.value.toString() + valorDer.value.toString()),type:tipo.STRING};
                else if(dominante == tipo.NULL) return {value: null, type:tipo.NULL} 
            }  

            
            case operador.MULTIPLICACION:{

                dominante = multi_division_Dominante[valorIz.type][valorDer.type];

                if(dominante == tipo.INTEGER ) return {value:(valorIz.value * valorDer.value),type:tipo.INTEGER};
                else if(dominante == tipo.FLOAT ) return {value:(valorIz.value * valorDer.value),type:tipo.FLOAT};
                else if(dominante == tipo.NULL) return {value: null, type:tipo.NULL} 
            } 
            
            
            case operador.DIVISION:
                {

                    dominante = multi_division_Dominante[valorIz.type][valorDer.type];
                    if(dominante == tipo.NULL) return {value: null, type:tipo.NULL} 
                    
                    if(!valorDer.value || valorDer.value == 0 )  return {value: null, type:tipo.NULL} 

                    if(dominante == tipo.INTEGER ) return {value:(valorIz.value / valorDer.value),type:tipo.INTEGER};
                    else if(dominante == tipo.FLOAT ) return {value:(valorIz.value / valorDer.value),type:tipo.FLOAT};
                    
                }


        }

        return null;
    }

    
}