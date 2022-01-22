"use strict";
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
const express_1 = __importDefault(require("express"));
const inicio_1 = require("../controllers/inicio");
const rutas = express_1.default.Router();
rutas.get("/", (request, response) => {
    return response.status(200).json({
        hola: "hola mundo"
    });
});
rutas.post("/prueba", function (request, response) {
    (0, inicio_1.probar)(request, response);
});
module.exports = rutas;
//# sourceMappingURL=route.js.map