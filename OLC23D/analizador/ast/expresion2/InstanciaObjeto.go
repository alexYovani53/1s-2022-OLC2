package expresion2

import (
	"OLC2/analizador"
	"OLC2/analizador/ast/definicion"
	"OLC2/analizador/entorno"
	"OLC2/analizador/entorno/Simbolos"
	"fmt"
	"reflect"
)

type InstanciaObjeto struct {
	Id string
}

func NewInstanciaObjeto(id string) InstanciaObjeto {
	return InstanciaObjeto{Id: id}
}

func (this InstanciaObjeto) Obtener3D(ent *entorno.Entorno) entorno.Result3D {

	RESULTADO_FINAL := entorno.Result3D{}

	existe := ent.ExisteClase(this.Id)

	if !existe {
		return entorno.Result3D{Tipo: entorno.NULL}
	}

	clasePlantilla := ent.ObtenerClase(this.Id).(*entorno.Clase)

	RESULTADO_FINAL.Codigo += fmt.Sprintf("/* **** VAMOS A DECLARAR UN OBJETO *************/\n")

	direccionStack := analizador.GeneradorGlobal.ObtenerTemporal()
	direccionHeap := analizador.GeneradorGlobal.ObtenerTemporal()

	tamanioEntornoActual := ent.Tamanio

	tamanioClase := this.tamano(clasePlantilla)

	RESULTADO_FINAL.Codigo += fmt.Sprintf("%s = HP; /* DIRECCIÓN INICIAL EN HEAP*/ \n ", direccionHeap)
	RESULTADO_FINAL.Codigo += fmt.Sprintf("HP = HP + %s; \n", tamanioClase)

	RESULTADO_FINAL.Codigo += fmt.Sprintf("%s = SP + %d; /* DIRECCION EN STACK*/\n", direccionStack, tamanioEntornoActual)
	RESULTADO_FINAL.Codigo += fmt.Sprintf("Stack[(int) %s] = %s  ;\n ", direccionStack, direccionHeap)

	ENTORNO_OBJETO := entorno.NewEntorno("OBJETO", ent)

	for i := 0; i < clasePlantilla.Instrucciones.Len(); i++ {

		INSTRU_PIVOTE := clasePlantilla.Instrucciones.GetValue(i)

		if reflect.TypeOf(INSTRU_PIVOTE) == reflect.TypeOf(Simbolos.Funcion{}) {
			//
		} else {

			DECLAR_ := INSTRU_PIVOTE.(*definicion.Declaracion)

			DECLAR_.CambiarEntorno = true
			DECLAR_.TemporalEntorno = direccionHeap

			RESULTADO_FINAL.Codigo += DECLAR_.Get3D(&ENTORNO_OBJETO)

		}

	}

	posicionRelativa := tamanioEntornoActual

	OBJETO_SIMBOLO := Simbolos.NewObjecto("", ENTORNO_OBJETO, clasePlantilla.Id)
	OBJETO_SIMBOLO.Direccion = posicionRelativa

	RESULTADO_FINAL.Tipo = entorno.OBJETO
	RESULTADO_FINAL.Valor = OBJETO_SIMBOLO

	return RESULTADO_FINAL
}

func (this InstanciaObjeto) tamano(clase *entorno.Clase) string {

	tamano := 0

	for i := 0; i < clase.Instrucciones.Len(); i++ {
		if reflect.TypeOf(clase.Instrucciones.GetValue(i)) != reflect.TypeOf(Simbolos.Funcion{}) {
			tamano = tamano + 1
		}
	}

	return fmt.Sprintf("%d", tamano)
}

func (this InstanciaObjeto) Obtener3DRef(ent *entorno.Entorno) entorno.Result3D {
	return entorno.Result3D{}
}
