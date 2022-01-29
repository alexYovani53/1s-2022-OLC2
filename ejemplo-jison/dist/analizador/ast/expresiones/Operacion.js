"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.Operacion = exports.operador = void 0;
const Expresion_1 = require("../abstract/Expresion");
const Retorno_1 = require("../abstract/Retorno");
var operador;
(function (operador) {
    operador[operador["MAS"] = 0] = "MAS";
    operador[operador["MENOS"] = 1] = "MENOS";
    operador[operador["MULTIPLICACION"] = 2] = "MULTIPLICACION";
    operador[operador["DIVISION"] = 3] = "DIVISION";
})(operador = exports.operador || (exports.operador = {}));
/**
 *     operando  int    float    string   boolean    null
 * int          [      |       |         |         |     ]
 * float        [      |       |         |         |     ]
 * string       [      |       |         |         |     ]
 * boolean      [      |       |         |         |     ]
 * null         [      |       |         |         |     ]
 */
const sumaDominante = [
    [Retorno_1.tipo.INTEGER, Retorno_1.tipo.FLOAT, Retorno_1.tipo.STRING, Retorno_1.tipo.NULL, Retorno_1.tipo.NULL],
    [Retorno_1.tipo.FLOAT, Retorno_1.tipo.FLOAT, Retorno_1.tipo.STRING, Retorno_1.tipo.NULL, Retorno_1.tipo.NULL],
    [Retorno_1.tipo.STRING, Retorno_1.tipo.STRING, Retorno_1.tipo.STRING, Retorno_1.tipo.STRING, Retorno_1.tipo.NULL],
    [Retorno_1.tipo.NULL, Retorno_1.tipo.NULL, Retorno_1.tipo.STRING, Retorno_1.tipo.NULL, Retorno_1.tipo.NULL],
    [Retorno_1.tipo.NULL, Retorno_1.tipo.NULL, Retorno_1.tipo.NULL, Retorno_1.tipo.NULL, Retorno_1.tipo.NULL],
];
const multi_division_Dominante = [
    [Retorno_1.tipo.INTEGER, Retorno_1.tipo.FLOAT, Retorno_1.tipo.NULL, Retorno_1.tipo.NULL, Retorno_1.tipo.NULL],
    [Retorno_1.tipo.FLOAT, Retorno_1.tipo.FLOAT, Retorno_1.tipo.NULL, Retorno_1.tipo.NULL, Retorno_1.tipo.NULL],
    [Retorno_1.tipo.NULL, Retorno_1.tipo.NULL, Retorno_1.tipo.NULL, Retorno_1.tipo.NULL, Retorno_1.tipo.NULL],
    [Retorno_1.tipo.NULL, Retorno_1.tipo.NULL, Retorno_1.tipo.NULL, Retorno_1.tipo.NULL, Retorno_1.tipo.NULL],
    [Retorno_1.tipo.NULL, Retorno_1.tipo.NULL, Retorno_1.tipo.NULL, Retorno_1.tipo.NULL, Retorno_1.tipo.NULL],
];
const restaDominante = [
    [Retorno_1.tipo.INTEGER, Retorno_1.tipo.FLOAT, Retorno_1.tipo.NULL, Retorno_1.tipo.NULL, Retorno_1.tipo.NULL],
    [Retorno_1.tipo.FLOAT, Retorno_1.tipo.FLOAT, Retorno_1.tipo.NULL, Retorno_1.tipo.NULL, Retorno_1.tipo.NULL],
    [Retorno_1.tipo.NULL, Retorno_1.tipo.NULL, Retorno_1.tipo.NULL, Retorno_1.tipo.NULL, Retorno_1.tipo.NULL],
    [Retorno_1.tipo.NULL, Retorno_1.tipo.NULL, Retorno_1.tipo.NULL, Retorno_1.tipo.NULL, Retorno_1.tipo.NULL],
    [Retorno_1.tipo.NULL, Retorno_1.tipo.NULL, Retorno_1.tipo.NULL, Retorno_1.tipo.NULL, Retorno_1.tipo.NULL],
];
class Operacion extends Expresion_1.Expresion {
    constructor(opIz, operador, opDer, unario, linea, columna) {
        super(linea, columna);
        this.operandoIz = opIz;
        this.operandoDe = opDer;
        this.operador = operador;
        this.unario = unario;
    }
    getValorImplicito() {
        let valorIz;
        let valorDer;
        let valorUn;
        if (this.unario) {
            valorUn = this.operandoIz.getValorImplicito();
        }
        else {
            valorIz = this.operandoIz.getValorImplicito();
            valorDer = this.operandoDe.getValorImplicito();
        }
        let dominante;
        switch (this.operador) {
            case operador.MAS: {
                dominante = sumaDominante[valorIz.type][valorDer.type];
                if (dominante == Retorno_1.tipo.INTEGER)
                    return { value: (valorIz.value + valorDer.value), type: Retorno_1.tipo.INTEGER };
                else if (dominante == Retorno_1.tipo.FLOAT)
                    return { value: (valorIz.value + valorDer.value), type: Retorno_1.tipo.FLOAT };
                else if (dominante == Retorno_1.tipo.STRING)
                    return { value: (valorIz.value.toString() + valorDer.value.toString()), type: Retorno_1.tipo.STRING };
                else if (dominante == Retorno_1.tipo.NULL)
                    return { value: null, type: Retorno_1.tipo.NULL };
            }
            case operador.MULTIPLICACION: {
                dominante = multi_division_Dominante[valorIz.type][valorDer.type];
                if (dominante == Retorno_1.tipo.INTEGER)
                    return { value: (valorIz.value * valorDer.value), type: Retorno_1.tipo.INTEGER };
                else if (dominante == Retorno_1.tipo.FLOAT)
                    return { value: (valorIz.value * valorDer.value), type: Retorno_1.tipo.FLOAT };
                else if (dominante == Retorno_1.tipo.NULL)
                    return { value: null, type: Retorno_1.tipo.NULL };
            }
            case operador.DIVISION:
                {
                    dominante = multi_division_Dominante[valorIz.type][valorDer.type];
                    if (dominante == Retorno_1.tipo.NULL)
                        return { value: null, type: Retorno_1.tipo.NULL };
                    if (!valorDer.value || valorDer.value == 0)
                        return { value: null, type: Retorno_1.tipo.NULL };
                    if (dominante == Retorno_1.tipo.INTEGER)
                        return { value: (valorIz.value / valorDer.value), type: Retorno_1.tipo.INTEGER };
                    else if (dominante == Retorno_1.tipo.FLOAT)
                        return { value: (valorIz.value / valorDer.value), type: Retorno_1.tipo.FLOAT };
                }
        }
        return null;
    }
}
exports.Operacion = Operacion;
//# sourceMappingURL=Operacion.js.map