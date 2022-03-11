package definicion

import (
	"OLC2/analizador/ast/interfaces"
	"OLC2/analizador/entorno"
)

type Asignacion struct {
	ID    string
	Expr  interfaces.Expresion
	Valor interface{}
}

func NewAsignacion(ID string, expr interfaces.Expresion) Asignacion {
	return Asignacion{ID: ID, Expr: expr}
}

func (a Asignacion) Ejecutar(ent entorno.Entorno) interface{} {

	existesVariable := ent.ExisteSimbolo(a.ID)

	valorExpr := a.Expr.ObtenerValor(ent)

	if !existesVariable {
		return nil
	}

	variable := ent.ObtenerSimbolo(a.ID).(entorno.Simbolo)

	if variable.EsReferencia {

		variable.Valor = valorExpr.Valor
		ent.CambiarValor(variable.Identificador, variable)

	} else {

		variable.Valor = valorExpr.Valor
		ent.CambiarValor(variable.Identificador, variable)
	}

	return nil
}
