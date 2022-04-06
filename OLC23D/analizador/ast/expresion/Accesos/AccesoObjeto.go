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
	ListaAccesos *arrayList.List
}

func NewAccesoObjeto(listaAccesos *arrayList.List) AccesoObjeto {
	return AccesoObjeto{ListaAccesos: listaAccesos}
}

func (a AccesoObjeto) Obtener3D(ent *entorno.Entorno) entorno.Result3D {

	EXPR_INICIAL := a.ListaAccesos.GetValue(0)

	if (reflect.TypeOf(EXPR_INICIAL) == reflect.TypeOf(expresion.Identificador{})) {

		ID := EXPR_INICIAL.(expresion.Identificador).Identificador

		if !ent.ExisteSimbolo(ID) {
			fmt.Printf("No existe OBJETOOOOOO")
			return entorno.Result3D{Tipo: entorno.NULL}
		}

		OBJETO := ent.ObtenerSimbolo(ID).(entorno.SimboloAbstracto)

		if OBJETO.GetTipo() == entorno.OBJETO {

			nuevaLista := a.ListaAccesos.Clone()
			nuevaLista.RemoveAtIndex(0)

			VALOR := OBJETO.(Simbolos.Objecto)
			RESULTADO_FINAL := entorno.Result3D{}
			DIR_STACK := analizador.GeneradorGlobal.ObtenerTemporal()
			DIR_HEAP := analizador.GeneradorGlobal.ObtenerTemporal()

			RESULTADO_FINAL.Codigo += "/* ACCEDIENDO A UN OBJETO */ \n"
			RESULTADO_FINAL.Codigo += fmt.Sprintf("%s = SP + %d;\n", DIR_STACK, VALOR.Direccion)
			RESULTADO_FINAL.Codigo += fmt.Sprintf("%s = Stack[(int)%s];\n", DIR_HEAP, DIR_STACK)

			DATO_ENCONTRADO := a.ObtenerValorRecursivo(ent, nuevaLista, VALOR, DIR_HEAP)

			RESULTADO_FINAL.Codigo += DATO_ENCONTRADO.Codigo
			RESULTADO_FINAL.Tipo = DATO_ENCONTRADO.Tipo
			RESULTADO_FINAL.Temporal = DATO_ENCONTRADO.Temporal

			return RESULTADO_FINAL

		}

	}

	return entorno.Result3D{Tipo: entorno.NULL}

}

func (a AccesoObjeto) ObtenerValorRecursivo(ent *entorno.Entorno, ListaAccesos *arrayList.List, OBJETO Simbolos.Objecto, DIR_HEAP string) entorno.Result3D {

	EXPR_INICIAL := ListaAccesos.GetValue(0)

	if (reflect.TypeOf(EXPR_INICIAL) == reflect.TypeOf(expresion.Identificador{})) {

		ID := EXPR_INICIAL.(expresion.Identificador).Identificador

		if !OBJETO.EntornoPropio.ExisteSimbolo(ID) {
			fmt.Printf("No existe OBJETOOOOOO")
			return entorno.Result3D{Tipo: entorno.NULL}
		}

		simb := OBJETO.EntornoPropio.ObtenerSimbolo(ID)
		simbolo := simb.(entorno.Simbolo)

		RESULTADO_FINAL := entorno.Result3D{}
		POS_HEAP := analizador.GeneradorGlobal.ObtenerTemporal()
		TEMP := analizador.GeneradorGlobal.ObtenerTemporal()

		RESULTADO_FINAL.Codigo += fmt.Sprintf("%s = %s + %d;\n", POS_HEAP, DIR_HEAP, simbolo.Direccion)
		RESULTADO_FINAL.Codigo += fmt.Sprintf("%s = Heap[(int) %s];\n", TEMP, POS_HEAP)

		RESULTADO_FINAL.Tipo = simbolo.Tipo
		RESULTADO_FINAL.Temporal = TEMP

		return RESULTADO_FINAL
	}
	return entorno.Result3D{}
}
