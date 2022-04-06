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

	existeClase := ent.ExisteClase(this.Id)

	if !existeClase {
		return entorno.Result3D{}
	}

	clasePlantilla := ent.ObtenerClase(this.Id).(*entorno.Clase)

	RESULTADO_FINAL.Codigo += "/******************************* DECLARANDO OBJETO ************************/ \n"

	direccionStack := analizador.GeneradorGlobal.ObtenerTemporal()
	direccionHeap := analizador.GeneradorGlobal.ObtenerTemporal()

	tamano := this.tamano(clasePlantilla)

	RESULTADO_FINAL.Codigo += fmt.Sprintf("%s = HP; /* DIRECCIÓN INICIAL DEL OBJETO*/\n", direccionHeap)
	RESULTADO_FINAL.Codigo += fmt.Sprintf("HP = HP + %s; /* APARTAR TAMAÑO DEL STRUCT */ \n", tamano)

	posicionRelativa := ent.Tamanio
	RESULTADO_FINAL.Codigo += fmt.Sprintf("%s = SP + %d;\n", direccionStack, posicionRelativa)
	RESULTADO_FINAL.Codigo += fmt.Sprintf("Heap[(int)%s] = %s;\n", direccionStack, direccionHeap)

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

				decl := r.(*definicion.Declaracion)
				decl.CambiarEntorno = true
				decl.TemporalEntorno = direccionHeap
				RESULTADO_FINAL.Codigo += decl.Get3D(&ENTORNO_OBJETO)
			}

		}

	}

	OBJETO_SIMBOLO := Simbolos.NewObjecto("", ENTORNO_OBJETO, this.Id)
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
