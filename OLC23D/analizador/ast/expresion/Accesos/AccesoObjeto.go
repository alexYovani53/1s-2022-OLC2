package Accesos

import (
	"OLC2/analizador"
	"OLC2/analizador/ast/expresion"
	"OLC2/analizador/entorno"
	"OLC2/analizador/entorno/Simbolos"
	"fmt"
	arrayList "github.com/colegno/arraylist"
	"reflect"
)

type AccesoObjeto struct {
	ListaAccesos      *arrayList.List
	ObtenerReferencia bool
}

func NewAccesoObjeto(listaAccesos *arrayList.List) AccesoObjeto {
	return AccesoObjeto{ListaAccesos: listaAccesos}
}

func (a AccesoObjeto) Obtener3D(ent *entorno.Entorno) entorno.Result3D {

	EXPR_INICIAL := a.ListaAccesos.GetValue(0)

	if reflect.TypeOf(EXPR_INICIAL) != reflect.TypeOf(expresion.Identificador{}) {
		return entorno.Result3D{Tipo: entorno.NULL}
	}

	ID := EXPR_INICIAL.(expresion.Identificador)
	existeEnEntorno := ent.ExisteSimbolo(ID.Identificador)

	if !existeEnEntorno {
		return entorno.Result3D{Tipo: entorno.NULL}
	}

	OBJETO := ent.ObtenerSimbolo(ID.Identificador).(entorno.SimboloAbstracto)

	if OBJETO.GetTipo() == entorno.OBJETO {

		nuevaLista := a.ListaAccesos.Clone()
		nuevaLista.RemoveAtIndex(0)

		VALOR := OBJETO.(Simbolos.Objecto)

		RESULTADO_FINAL := entorno.Result3D{}

		DIRECCION_STACK := analizador.GeneradorGlobal.ObtenerTemporal()
		DIRECCION_Heap := analizador.GeneradorGlobal.ObtenerTemporal()

		RESULTADO_FINAL.Codigo += "/* ACCEDIENDO A UN OBJETO */ \n"

		RESULTADO_FINAL.Codigo += fmt.Sprintf("%s = SP + %d;  /* CAPTURANDO DIRECCION EN STACK*/ \n", DIRECCION_STACK, VALOR.Direccion)
		RESULTADO_FINAL.Codigo += fmt.Sprintf("%s = Stack[(int) %s]; \n", DIRECCION_Heap, DIRECCION_STACK)

		BUSQUEDA := a.ObtenerValorRecursivo(ent, nuevaLista, VALOR, DIRECCION_Heap)

		RESULTADO_FINAL.Codigo += BUSQUEDA.Codigo
		RESULTADO_FINAL.Tipo = BUSQUEDA.Tipo
		RESULTADO_FINAL.Temporal = BUSQUEDA.Temporal
		RESULTADO_FINAL.ValorEnHeap = BUSQUEDA.ValorEnHeap

		return RESULTADO_FINAL
	}

	return entorno.Result3D{Tipo: entorno.NULL}
}

func (a AccesoObjeto) ObtenerValorRecursivo(ent *entorno.Entorno, ListaAccesos *arrayList.List, OBJETO Simbolos.Objecto, DIR_HEAP string) entorno.Result3D {

	EXPR_INI := ListaAccesos.GetValue(0)

	if reflect.TypeOf(EXPR_INI) != reflect.TypeOf(expresion.Identificador{}) {
		return entorno.Result3D{Tipo: entorno.NULL}
	}

	ID := EXPR_INI.(expresion.Identificador)

	existeVariable := OBJETO.EntornoPropio.ExisteSimbolo(ID.Identificador)
	if !existeVariable {
		return entorno.Result3D{Tipo: entorno.NULL}
	}

	Simbolo_ := OBJETO.EntornoPropio.ObtenerSimbolo(ID.Identificador).(entorno.Simbolo)

	RESULTADO_FINAL := entorno.Result3D{}
	POS_HEAP := analizador.GeneradorGlobal.ObtenerTemporal()
	TEMPORAL := analizador.GeneradorGlobal.ObtenerTemporal()

	RESULTADO_FINAL.Codigo += fmt.Sprintf("%s = %s  + %d;\n", POS_HEAP, DIR_HEAP, Simbolo_.Direccion)

	if a.ObtenerReferencia {
		RESULTADO_FINAL.Temporal = POS_HEAP
		RESULTADO_FINAL.ValorEnHeap = true
	} else {

		RESULTADO_FINAL.Codigo += fmt.Sprintf("%s = Heap[(int) %s];/*CAPTURANDO VARIABLE %s */;\n ", TEMPORAL, POS_HEAP, ID.Identificador)
		RESULTADO_FINAL.Temporal = TEMPORAL
	}

	RESULTADO_FINAL.Tipo = Simbolo_.Tipo

	return RESULTADO_FINAL
}

func (this AccesoObjeto) Obtener3DRef(ent *entorno.Entorno) entorno.Result3D {
	this.ObtenerReferencia = true

	return this.Obtener3D(ent)
}
