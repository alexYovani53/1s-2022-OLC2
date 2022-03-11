package Accesos

import (
	"OLC2/analizador/ast/expresion"
	"OLC2/analizador/entorno"
	"OLC2/analizador/entorno/Simbolos"
	"fmt"
	arrayList "github.com/colegno/arraylist"
	"reflect"
)

type AccesoObjeto struct {
	ListaAccesos *arrayList.List
}

func NewAccesoObjeto(listaAccesos *arrayList.List) AccesoObjeto {
	return AccesoObjeto{ListaAccesos: listaAccesos}
}

func (a AccesoObjeto) ObtenerValor(ent entorno.Entorno) entorno.ValorType {
	//TODO implement me

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

			return a.ObtenerValorRecursivo(ent, nuevaLista, VALOR)

		}

	}

	return entorno.ValorType{Valor: nil, Tipo: entorno.NULL}
}

func (a AccesoObjeto) ObtenerValorRecursivo(ent entorno.Entorno, ListaAccesos *arrayList.List, OBJETO Simbolos.Objecto) entorno.ValorType {

	EXPR_INICIAL := ListaAccesos.GetValue(0)

	if (reflect.TypeOf(EXPR_INICIAL) == reflect.TypeOf(expresion.Identificador{})) {

		ID := EXPR_INICIAL.(expresion.Identificador).Identificador

		if !OBJETO.EntornoPropio.ExisteSimbolo(ID) {
			fmt.Printf("No existe OBJETOOOOOO")
			return entorno.ValorType{Valor: nil, Tipo: entorno.NULL}
		}

		simbolo := OBJETO.EntornoPropio.ObtenerSimbolo(ID)

		return entorno.ValorType{Valor: simbolo.(entorno.Simbolo).Valor, Tipo: simbolo.(entorno.Simbolo).Tipo}

	}

	return entorno.ValorType{Valor: nil, Tipo: entorno.NULL}
}
