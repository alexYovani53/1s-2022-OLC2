
import { Ast } from "analizador/ast/Ast";
import { Entorno } from "../analizador/entorno/Entorno";
import { errores, limpiar } from "../analizador/Errores";

const analizador = require("../analizador/gramatica/gramatica.js");


export  async function probar (request,response){

    limpiar();

    let entorno:Entorno =  new Entorno(null,"hola");

    let result:Ast = analizador.parse(request.body.text);
    let val = "";
    for (const iterator of result.instrucciones) {
        val += iterator.ejecutar(entorno);
    }

    response.status(200).json({
        errores,
        result,
        val,
        estado:"exito"
    })
}