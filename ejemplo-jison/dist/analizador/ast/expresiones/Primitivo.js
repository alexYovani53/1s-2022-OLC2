"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.Primitivo = void 0;
const Expresion_1 = require("../abstract/Expresion");
class Primitivo extends Expresion_1.Expresion {
    constructor(valor, tipo, linea, columna) {
        super(linea, columna);
        this.valor = valor;
        this.tipo = tipo;
    }
    getValorImplicido(ent) {
        return { value: this.valor, type: this.tipo };
    }
}
exports.Primitivo = Primitivo;
//# sourceMappingURL=Primitivo.js.map