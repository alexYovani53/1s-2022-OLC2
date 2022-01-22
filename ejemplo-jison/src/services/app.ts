import {createServer} from "http";
import express from "express";
import bodyParser from "body-parser";
import cors from "cors";
import rutas from "../routers/route"

const corsOptions = { origin: true, optionsSuccessStatus: 200 };

const servidor = express();
const httpServer = createServer(servidor);
servidor.use(cors(corsOptions));
servidor.use(bodyParser.json({ limit: '10mb' }));
servidor.use(bodyParser.urlencoded({ limit: '10mb', extended: true }));

servidor.use('/',rutas);


export const inicializar = ()=>{
    return new Promise<void>((resolve, reject)=>{
             
        httpServer.listen(3000)
        .on('listening',()=>{
            console.log("Servidor ejecutandose");
            resolve();
        })
        .on("error",(err)=>{
            console.log("Falla en el servidor");
            reject(err)
        })

    });


}


export function close(){
    return new Promise<void>((resolve,reject)=>{
        httpServer.close(err=>{
            if(err) return reject(err);
            return resolve();
        })
    })
}

