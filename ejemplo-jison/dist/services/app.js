"use strict";
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
exports.close = exports.inicializar = void 0;
const http_1 = require("http");
const express_1 = __importDefault(require("express"));
const body_parser_1 = __importDefault(require("body-parser"));
const cors_1 = __importDefault(require("cors"));
const route_1 = __importDefault(require("../routers/route"));
const corsOptions = { origin: true, optionsSuccessStatus: 200 };
const servidor = (0, express_1.default)();
const httpServer = (0, http_1.createServer)(servidor);
servidor.use((0, cors_1.default)(corsOptions));
servidor.use(body_parser_1.default.json({ limit: '10mb' }));
servidor.use(body_parser_1.default.urlencoded({ limit: '10mb', extended: true }));
servidor.use('/', route_1.default);
const inicializar = () => {
    return new Promise((resolve, reject) => {
        httpServer.listen(3000)
            .on('listening', () => {
            console.log("Servidor ejecutandose");
            resolve();
        })
            .on("error", (err) => {
            console.log("Falla en el servidor");
            reject(err);
        });
    });
};
exports.inicializar = inicializar;
function close() {
    return new Promise((resolve, reject) => {
        httpServer.close(err => {
            if (err)
                return reject(err);
            return resolve();
        });
    });
}
exports.close = close;
//# sourceMappingURL=app.js.map