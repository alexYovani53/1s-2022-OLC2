package Simbolos

import (
	"OLC2/analizador"
	"OLC2/analizador/ast/interfaces"
	"OLC2/analizador/entorno"
	"github.com/colegno/arraylist"
)

var retornoVal = [6][56]entorno.TipoDato{
	{entorno.INTEGER, entorno.FLOAT, entorno.NULL, entorno.NULL, entorno.NULL, entorno.NULL},
	{entorno.FLOAT, entorno.FLOAT, entorno.NULL, entorno.NULL, entorno.NULL, entorno.NULL},
	{entorno.NULL, entorno.NULL, entorno.STRING, entorno.NULL, entorno.NULL, entorno.NULL},
	{entorno.NULL, entorno.NULL, entorno.NULL, entorno.BOOLEAN, entorno.NULL, entorno.NULL},
	{entorno.NULL, entorno.NULL, entorno.NULL, entorno.NULL, entorno.NULL, entorno.NULL},
	{entorno.NULL, entorno.NULL, entorno.NULL, entorno.NULL, entorno.NULL, entorno.VOID},
}

type Funcion struct {
	entorno.Simbolo
	ListaParamsDecl    *arraylist.List
	ListaInstrucciones *arraylist.List
}

func NewFuncion(nombre string, listaParams *arraylist.List, listaInstrucciones *arraylist.List, tipo entorno.TipoDato) Funcion {
	funcSimbolo := entorno.NewSimboloFuncion(0, 0, nombre, tipo, listaParams)

	return Funcion{
		ListaInstrucciones: listaInstrucciones,
		ListaParamsDecl:    listaParams,
		Simbolo:            funcSimbolo,
	}
}

func (f Funcion) EjecutarParametros(ent *entorno.Entorno, expresiones *arraylist.List, entRef *entorno.Entorno) bool {

	return true
}

func (f Funcion) Get3D(ent *entorno.Entorno) string {

	etiquetaRetonro := analizador.GeneradorGlobal.ObtenerEtiqueta()

	ENTORNO_FUNCION := entorno.NewEntorno("Funcion", ent)

	codigoFuncion := ""

	codigoFuncion += "void " + f.Identificador + " {\n\n"

	for i := 0; i < f.ListaInstrucciones.Len(); i++ {

		INSTR := f.ListaInstrucciones.GetValue(i)

		codigoINSTR := INSTR.(interfaces.Instruccion).Get3D(&ENTORNO_FUNCION)

		codigoFuncion += analizador.GeneradorGlobal.Tabular(codigoINSTR)
	}

	codigoFuncion += analizador.GeneradorGlobal.TabularLinea(etiquetaRetonro+"\n", 1)
	codigoFuncion += analizador.GeneradorGlobal.TabularLinea("return\n", 1)

	codigoFuncion += "}"

	return codigoFuncion
}

/*
fn is_even(a:u8) -> bool {
    if a % 2 == 0 {
        return true
    }
    return false // el valor de retorno en todos los otros casos
}

*/
