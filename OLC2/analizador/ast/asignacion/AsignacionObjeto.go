package asignacion

import (
	"OLC2/analizador/ast/expresion"
	"OLC2/analizador/ast/interfaces"
	"OLC2/analizador/entorno"
	"OLC2/analizador/entorno/Simbolos"
	"fmt"
	arrayList "github.com/colegno/arraylist"
	"reflect"
)

type AsignacionObjeto struct {
	ListaAccesos *arrayList.List
	Expr         interfaces.Expresion
	Valor        interface{}
}

func NewAsignacionObjeto(listaAccesos *arrayList.List, expr interfaces.Expresion) AsignacionObjeto {
	return AsignacionObjeto{ListaAccesos: listaAccesos, Expr: expr}
}

func NewAsignacionObjetoValor(listaAccesos *arrayList.List, valor interface{}) AsignacionObjeto {
	return AsignacionObjeto{ListaAccesos: listaAccesos, Valor: valor}
}

func (a AsignacionObjeto) Ejecutar(ent entorno.Entorno) interface{} {

	// ID .   ID
	EXPR_INICIAL := a.ListaAccesos.GetValue(0)

	if (reflect.TypeOf(EXPR_INICIAL) == reflect.TypeOf(expresion.Identificador{})) {

		ID := EXPR_INICIAL.(expresion.Identificador).Identificador

		if !ent.ExisteSimbolo(ID) {
			fmt.Printf("No existe OBJETOOOOOO")
			return entorno.ValorType{Valor: nil, Tipo: entorno.NULL}
		}

		OBJETO := ent.ObtenerSimbolo(ID).(entorno.SimboloAbstracto)

		if OBJETO.GetTipo() == entorno.OBJETO {

			VALOR := OBJETO.(Simbolos.Objecto)

			nuevaLista := a.ListaAccesos.Clone()
			nuevaLista.RemoveAtIndex(0)

			return a.CambiarValorRecursivo(ent, nuevaLista, VALOR)

		}

	}

	return nil
}

func (a AsignacionObjeto) CambiarValorRecursivo(ent entorno.Entorno, ListaAccesos *arrayList.List, OBJETO Simbolos.Objecto) entorno.ValorType {

	// PERSONA.CASA.TELEVISOR

	EXPR_INICIAL := ListaAccesos.GetValue(0)
	nuevaLista := a.ListaAccesos.Clone()
	nuevaLista.RemoveAtIndex(0)

	if (reflect.TypeOf(EXPR_INICIAL) == reflect.TypeOf(expresion.Identificador{})) {

		ID := EXPR_INICIAL.(expresion.Identificador).Identificador

		if !OBJETO.EntornoPropio.ExisteSimbolo(ID) {
			fmt.Printf("No existe OBJETOOOOOO")
			return entorno.ValorType{Valor: nil, Tipo: entorno.NULL}
		}

		simbolo := OBJETO.EntornoPropio.ObtenerSimbolo(ID)

		if reflect.TypeOf(simbolo) != reflect.TypeOf(entorno.Simbolo{}) {
			//return entorno.ValorType{Valor: -1, Tipo: entorno.NULL}
			SUB_OBJETO := simbolo.(Simbolos.Objecto)
			return a.CambiarValorRecursivo(ent, nuevaLista, SUB_OBJETO)
		}

		var expresionValor entorno.ValorType

		if a.Expr != nil {
			expresionValor = a.Expr.ObtenerValor(ent)
		} else {
			expresionValor.Valor = a.Valor
		}

		if a.Valor != nil {
			if expresionValor.Tipo != simbolo.(entorno.Simbolo).Tipo {
				return entorno.ValorType{Valor: nil, Tipo: entorno.NULL}
			}

			simb := simbolo.(entorno.Simbolo)
			simb.Valor = expresionValor.Valor
			OBJETO.EntornoPropio.CambiarValor(simb.Identificador, simb)

		}

		return entorno.ValorType{Valor: true, Tipo: simbolo.(entorno.Simbolo).Tipo}

	}

	return entorno.ValorType{Valor: nil, Tipo: entorno.NULL}
}
