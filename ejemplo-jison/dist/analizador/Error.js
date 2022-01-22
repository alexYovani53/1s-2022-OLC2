"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.Error_ = void 0;
class Error_ {
    constructor(linea, columna, tipo, mensaje) {
        this.linea = linea;
        this.columna = columna;
        this.tipo = tipo;
        this.mensaje = mensaje;
    }
    getLinea() { return this.linea; }
    getColumna() { return this.columna; }
    getTipo() { return this.tipo; }
    getMensaje() { return this.mensaje; }
}
exports.Error_ = Error_;
//# sourceMappingURL=Error.js.map