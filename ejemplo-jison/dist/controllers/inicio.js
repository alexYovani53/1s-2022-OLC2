"use strict";
var __awaiter = (this && this.__awaiter) || function (thisArg, _arguments, P, generator) {
    function adopt(value) { return value instanceof P ? value : new P(function (resolve) { resolve(value); }); }
    return new (P || (P = Promise))(function (resolve, reject) {
        function fulfilled(value) { try { step(generator.next(value)); } catch (e) { reject(e); } }
        function rejected(value) { try { step(generator["throw"](value)); } catch (e) { reject(e); } }
        function step(result) { result.done ? resolve(result.value) : adopt(result.value).then(fulfilled, rejected); }
        step((generator = generator.apply(thisArg, _arguments || [])).next());
    });
};
Object.defineProperty(exports, "__esModule", { value: true });
exports.probar = void 0;
const Errores_1 = require("../analizador/Errores");
const analizador = require("../analizador/gramatica/gramatica.js");
function probar(request, response) {
    return __awaiter(this, void 0, void 0, function* () {
        (0, Errores_1.limpiar)();
        let result = analizador.parse(request.body.text);
        let val = "";
        for (const iterator of result.instrucciones) {
            val += iterator.ejecutar();
        }
        response.status(200).json({
            errores: Errores_1.errores,
            result,
            val,
            estado: "exito"
        });
    });
}
exports.probar = probar;
//# sourceMappingURL=inicio.js.map