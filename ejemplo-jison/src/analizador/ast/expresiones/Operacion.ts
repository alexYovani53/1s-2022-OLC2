import { Error_ } from "analizador/Error";
import { errores } from "analizador/Errores";
import { Expresion } from "../abstract/Expresion";
import { Retorno, tipo } from "../abstract/Retorno";
import { Primitivo } from "./Primitivo";



export enum operador{
    MAS,
    MENOS, 
    MULTIPLICACION,
    DIVISION
}

export class Operacion extends Expresion{

    operandoIz: Expresion;
    operandoDe: Expresion;
    operandoUn: Expresion;
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

        switch(this.operador){

            case operador.MAS:{

                if(valorIz.type==tipo.INTEGER && valorDer.type == tipo.INTEGER) return {value:(valorIz.value + valorDer.value),type:tipo.INTEGER};
                if(valorIz.type==tipo.INTEGER && valorDer.type == tipo.FLOAT) return {value:(valorIz.value + valorDer.value),type:tipo.FLOAT};
                if(valorIz.type==tipo.INTEGER && valorDer.type == tipo.STRING) return {value:(valorIz.value + valorDer.value),type:tipo.STRING};
                if(valorIz.type==tipo.INTEGER && valorDer.type == tipo.BOOLEAN) {
                    errores.push(new Error_(1,2,"Semantico","No se puede operar Integer con boolean"));
                }

            }  

            
            case operador.MULTIPLICACION:{

                if(valorIz.type==tipo.INTEGER && valorDer.type == tipo.INTEGER) return {value:(valorIz.value * valorDer.value),type:tipo.INTEGER};
                if(valorIz.type==tipo.INTEGER && valorDer.type == tipo.FLOAT) return {value:(valorIz.value *valorDer.value),type:tipo.FLOAT};

            }  


        }

        return null;
    }

    
}