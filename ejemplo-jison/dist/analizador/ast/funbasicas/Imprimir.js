"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.Imprimir = void 0;
const Instruccion_1 = require("../abstract/Instruccion");
class Imprimir extends Instruccion_1.Instruccion {
    constructor(expresion, linea, columna) {
        super(linea, columna);
        this.expresion = expresion;
    }
    ejecutar() {
        let val = this.expresion.getValorImplicito();
        return val.value + "\n";
    }
}
exports.Imprimir = Imprimir;
//# sourceMappingURL=Imprimir.js.map