package entorno

import (
	"github.com/colegno/arraylist"
)

type Simbolo struct {
	Linea         int
	Columna       int
	Identificador string
	Valor         interface{}
	Tipo          TipoDato
	Constante     bool
	EsFuncion     bool
	ListaParams   *arraylist.List
}

/**
 * @comentario  Constructor que se utiliza solo para guardar un ID en un simbolo
 * @param       identificador  -> Identificador del Simbolo.
 */
/*

	int d =  x + y;

*/

func NewSimboloIdentificador(linea int, columna int, identificador string) *Simbolo {
	return &Simbolo{
		Linea:         linea,
		Columna:       columna,
		Identificador: identificador,
		Constante:     false,
		EsFuncion:     false,
		Valor:         nil,
	}
}

func NewSimboloIdentificadorValor(linea int, columna int, identificador string, valor interface{}, dato TipoDato) Simbolo {
	e := Simbolo{
		Linea:         linea,
		Columna:       columna,
		Identificador: identificador,
		Constante:     false,
		EsFuncion:     false,
		Valor:         valor,
		Tipo:          dato,
	}
	return e
}

func NewSimboloFuncion(linea int, columna int, identificador string, tipoRet TipoDato, listParametros *arraylist.List) Simbolo {
	e := Simbolo{
		Linea:         linea,
		Columna:       columna,
		Identificador: identificador,
		Constante:     false,
		EsFuncion:     true,
		Valor:         nil,
		Tipo:          tipoRet,
		ListaParams:   listParametros,
	}
	return e
}
func NewSimboloArreglo(linea int, columna int, identificador string, tipoDatos TipoDato) Simbolo {
	e := Simbolo{
		Linea:         linea,
		Columna:       columna,
		Identificador: identificador,
		Constante:     false,
		EsFuncion:     true,
		Valor:         nil,
		Tipo:          tipoDatos,
		ListaParams:   nil,
	}
	return e
}
