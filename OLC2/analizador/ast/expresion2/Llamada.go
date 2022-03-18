package expresion2

import (
	"OLC2/analizador/entorno"
	"OLC2/analizador/entorno/Simbolos"
	arrayList "github.com/colegno/arraylist"
	"reflect"
)

type Llamada struct {
	IdFuncion        string
	ListaExpresiones *arrayList.List
}

func NewLlamada(idFuncion string, ListaExpresiones *arrayList.List) Llamada {
	return Llamada{IdFuncion: idFuncion, ListaExpresiones: ListaExpresiones}
}

func (l Llamada) ObtenerValor(ent entorno.Entorno) entorno.ValorType {

	hayRuncion := ent.ExisteFuncion(l.IdFuncion)

	if !hayRuncion {
		return entorno.ValorType{Valor: -1, Tipo: entorno.NULL}
	}

	// ENTORNO DE FUNCION
	entFunc := entorno.NewEntorno("Funcion", &ent)

	funcion := ent.ObtenerFuncion(l.IdFuncion).(Simbolos.Funcion)
	clonFunc := Simbolos.NewFuncion(funcion.Identificador, funcion.ListaParamsDecl.Clone(), funcion.ListaInstrucciones.Clone(), funcion.Tipo)

	completo := clonFunc.EjecutarParametros(entFunc, l.ListaExpresiones, &ent)

	if !completo {
		return entorno.ValorType{Valor: -1, Tipo: entorno.NULL}
	}

	val := clonFunc.Ejecutar(entFunc)

	if (reflect.TypeOf(val) != reflect.TypeOf(entorno.ValorType{})) {
		return entorno.ValorType{}
	}

	return val.(entorno.ValorType)
}

func (l Llamada) Ejecutar(ent entorno.Entorno) interface{} {

	l.ObtenerValor(ent)

	return nil
}
