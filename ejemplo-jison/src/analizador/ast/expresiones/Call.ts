import { Expresion } from "../abstract/Expresion";
import { Retorno } from "../abstract/Retorno";


export class Call extends Expresion{
    
    public getValorImplicito(): Retorno {
        throw new Error("Method not implemented.");
    }
    
}