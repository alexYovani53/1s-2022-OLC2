"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.Entorno = void 0;
class Entorno {
    constructor(entornoAnterio, nombre) {
        this.tabla = new Map();
        this.ent_anterior = entornoAnterio;
        this.nombre = nombre;
    }
}
exports.Entorno = Entorno;
//# sourceMappingURL=Entorno.js.map