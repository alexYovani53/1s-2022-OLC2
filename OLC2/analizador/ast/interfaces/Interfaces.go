package interfaces

import (
	"OLC2/analizador/entorno"
)

type Expresion interface {
	ObtenerValor(ent entorno.Entorno) entorno.ValorType
}

type Instruccion interface {
	Ejecutar(ent entorno.Entorno) interface{}
}

// type Instruccion interface {

// 	ejecutar()
// }
