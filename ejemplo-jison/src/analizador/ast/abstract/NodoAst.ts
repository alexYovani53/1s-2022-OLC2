


export abstract class  NodoAst {

    constructor(private line:number, private columna:number) {
        this.line = line;
        this.columna = columna;
    }

    getLine(): number {
        return this.line;
    }
    
    getColumna():number{
        return this.columna;
    }
}