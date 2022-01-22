export class Error_ {

    constructor(private linea:number, private columna:number, private tipo:string, private mensaje:string){

    }

    getLinea()      {return this.linea}
    getColumna()    {return this.columna}
    getTipo()       {return this.tipo}
    getMensaje()    {return this.mensaje}

    

}