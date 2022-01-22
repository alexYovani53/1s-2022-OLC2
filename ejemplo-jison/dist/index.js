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
const app_1 = require("./services/app");
const iniciar = () => __awaiter(void 0, void 0, void 0, function* () {
    try {
        yield (0, app_1.inicializar)();
    }
    catch (err) {
        console.log(err);
        console.log("no se iniciaron correctaente los servicios");
    }
});
const apagar = () => __awaiter(void 0, void 0, void 0, function* () {
    let err;
    try {
        yield (0, app_1.close)();
    }
    catch (error) {
        console.log(error);
        err = error;
    }
    if (err)
        process.exit(1);
    process.exit(0);
});
iniciar();
process.on('SIGTERM', () => apagar());
process.on('SIGINT', () => apagar());
process.on('uncaughtException', (err) => {
    console.log(err);
    apagar();
});
//# sourceMappingURL=index.js.map