package funbasica

import (
	"OLC2/analizador/ast/interfaces"
	"OLC2/analizador/entorno"
)

type Imprimir struct {
	Expresiones interfaces.Expresion
}

func NewImprimir(val interfaces.Expresion) Imprimir {
	e := Imprimir{val}
	return e
}

func (p Imprimir) Ejecutar(entorno entorno.Entorno) interface{} {

	retornoExpr := p.Expresiones.ObtenerValor(entorno)

	return retornoExpr.Valor
}
