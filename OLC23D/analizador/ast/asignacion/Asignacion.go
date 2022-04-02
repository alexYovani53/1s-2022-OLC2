package asignacion

import (
	"OLC2/analizador/ast/interfaces"
	"OLC2/analizador/entorno"
)

type Asignacion struct {
	ID    string
	Expr  interfaces.Expresion
	Valor interface{} //asignacion con un valor objeto pasado desde alguna instruccion
}

func NewAsignacion(ID string, expr interfaces.Expresion) Asignacion {
	return Asignacion{ID: ID, Expr: expr, Valor: nil}
}

func NewAsignacionValor(ID string, valor interface{}) Asignacion {
	return Asignacion{ID: ID, Expr: nil, Valor: valor}
}

func (a Asignacion) Get3D(ent *entorno.Entorno) string {

	return ""
}
