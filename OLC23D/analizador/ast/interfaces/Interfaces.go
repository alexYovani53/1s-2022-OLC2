package interfaces

import (
	"OLC2/analizador/entorno"
)

type Expresion interface {
	Obtener3D(ent *entorno.Entorno) entorno.Result3D
	Obtener3DRef(ent *entorno.Entorno) entorno.Result3D
}

type Instruccion interface {
	Get3D(ent *entorno.Entorno) string
}

// type Instruccion interface {

// 	ejecutar()
// }
