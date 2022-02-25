package funbasica

import (
	analizador2 "OLC2/analizador"
	"OLC2/analizador/ast/interfaces"
	"OLC2/analizador/entorno"
	"fmt"
)

type Imprimir struct {
	Expresiones interfaces.Expresion
}

func NewImprimir(val interfaces.Expresion) Imprimir {
	e := Imprimir{val}
	return e
}

func (p Imprimir) Ejecutar(ent entorno.Entorno) interface{} {

	retornoExpr := p.Expresiones.ObtenerValor(ent)
	conSalto := fmt.Sprintf("%v", retornoExpr.Valor)
	conSalto = conSalto + "\n"
	analizador2.Consola += conSalto
	
	return nil
}
