import express from "express";
import { probar} from "../controllers/inicio";

const rutas = express.Router();


rutas.get("/",(request,response)=>{
    return response.status(200).json({
        hola:"hola mundo"
    })
});

rutas.post("/prueba",function (request,response){
    probar(request,response);
});

export = rutas;