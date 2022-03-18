package asignacion

import (
	"OLC2/analizador/ast/interfaces"
	"OLC2/analizador/entorno"
	"fmt"
	"reflect"
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

func (a Asignacion) Ejecutar(ent entorno.Entorno) interface{} {

	if !ent.ExisteSimbolo(a.ID) {
		return nil
	}

	simbolo := ent.ObtenerSimbolo(a.ID)

	if reflect.TypeOf(simbolo) == reflect.TypeOf(entorno.Simbolo{}) {

		simboloPrimitivo := simbolo.(entorno.Simbolo)

		var valorAsignacion entorno.ValorType

		if a.Expr != nil {
			valorAsignacion = a.Expr.ObtenerValor(ent)
		} else {
			valorAsignacion.Valor = a.Valor
		}

		if simboloPrimitivo.EsReferencia && a.Expr != nil {

			if valorAsignacion.Tipo != simboloPrimitivo.Tipo {
				fmt.Printf("Error de tipos \n")
				return nil
			}

			simboloPrimitivo.Valor = valorAsignacion.Valor
			ent.CambiarValor(simboloPrimitivo.Identificador, simboloPrimitivo)

			valRef := NewValorRefEjecutar(simboloPrimitivo.Referencia)
			valRef.Ejecutar(ent, valorAsignacion.Valor)

		} else if simboloPrimitivo.EsReferencia && a.Expr == nil {

			simboloPrimitivo.Valor = valorAsignacion.Valor
			ent.CambiarValor(simboloPrimitivo.Identificador, simboloPrimitivo)

			valRef := NewValorRefEjecutar(simboloPrimitivo.Referencia)
			valRef.Ejecutar(ent, valorAsignacion.Valor)

		} else if !simboloPrimitivo.EsReferencia && a.Expr != nil {

			if valorAsignacion.Tipo != simboloPrimitivo.Tipo {
				fmt.Printf("Error de tipos \n")
				return nil
			}

			simboloPrimitivo.Valor = valorAsignacion.Valor
			ent.CambiarValor(simboloPrimitivo.Identificador, simboloPrimitivo)

		} else if !simboloPrimitivo.EsReferencia && a.Expr == nil {

			simboloPrimitivo.Valor = valorAsignacion.Valor
			ent.CambiarValor(simboloPrimitivo.Identificador, simboloPrimitivo)
		}

	}

	return nil
}
