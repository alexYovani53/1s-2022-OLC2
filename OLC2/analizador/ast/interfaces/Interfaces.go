package interfaces

import (
	"OLC2/analizador/entorno"
)

type Expresion interface {
	ObtenerValor(entorno entorno.Entorno) entorno.RetornoType
}

type Instruccion interface {
	Ejecutar(entorno entorno.Entorno) interface{}
}

// type Instruccion interface {

// 	ejecutar()
// }
