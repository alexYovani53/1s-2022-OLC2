package expresion2

import (
	"OLC2/analizador/ast/interfaces"
	"OLC2/analizador/entorno"
	"OLC2/analizador/entorno/Simbolos"
	"reflect"
)

type InstanciaObjeto struct {
	Id string
}

func NewInstanciaObjeto(id string) InstanciaObjeto {
	return InstanciaObjeto{Id: id}
}

func (i InstanciaObjeto) ObtenerValor(ent entorno.Entorno) entorno.ValorType {

	existeClase := ent.ExisteClase(i.Id)

	if !existeClase {
		return entorno.ValorType{Valor: nil, Tipo: entorno.NULL}
	}

	clasePlantilla := ent.ObtenerClase(i.Id).(*entorno.Clase)

	ENTORNO_OBJETO := entorno.NewEntorno("OBJETO", nil)

	for i := 0; i < clasePlantilla.Instrucciones.Len(); i++ {

		r := clasePlantilla.Instrucciones.GetValue(i)

		if r != nil {
			if reflect.TypeOf(r) == reflect.TypeOf(Simbolos.Funcion{}) {
				func_ := r.(Simbolos.Funcion)

				if !ENTORNO_OBJETO.ExisteFuncion(func_.Identificador) {
					ENTORNO_OBJETO.AgregarFuncion(func_.Identificador, func_)

				} else {
					//ERROR
				}

			} else {
				r.(interfaces.Instruccion).Ejecutar(ENTORNO_OBJETO)
			}

		}

	}

	OBJETO_SIMBOLO := Simbolos.NewObjecto("", ENTORNO_OBJETO, i.Id)

	return entorno.ValorType{Valor: OBJETO_SIMBOLO, Tipo: entorno.OBJETO}
}
