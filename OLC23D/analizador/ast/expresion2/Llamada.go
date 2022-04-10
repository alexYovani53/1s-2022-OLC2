package expresion2

import (
	"OLC2/analizador"
	"OLC2/analizador/entorno"
	"OLC2/analizador/entorno/Simbolos"
	"fmt"
	arrayList "github.com/colegno/arraylist"
)

type Llamada struct {
	IdFuncion        string
	ListaExpresiones *arrayList.List
}

func NewLlamada(idFuncion string, ListaExpresiones *arrayList.List) Llamada {
	return Llamada{IdFuncion: idFuncion, ListaExpresiones: ListaExpresiones}
}

func (this Llamada) Obtener3D(ent *entorno.Entorno) entorno.Result3D {

	SALIDA_FINAL := entorno.Result3D{}

	hayFuncion := ent.ExisteFuncion(this.IdFuncion)
	if !hayFuncion {
		return entorno.Result3D{Tipo: entorno.NULL}
	}

	Entorno_fun := entorno.NewEntorno("Funcion", ent)

	Funcion := ent.ObtenerFuncion(this.IdFuncion).(Simbolos.Funcion)
	clonFunc := Simbolos.NewFuncion(Funcion.Identificador, Funcion.ListaParamsDecl.Clone(), Funcion.ListaInstrucciones.Clone(), Funcion.Tipo)

	this.ValidarFuncionGenerada(ent, &Funcion)

	Entorno_fun.Tamanio = 1
	// STACK_FUNCION [ --retorno--  , -- param1--, -- param2--, -- paramN--, -- VariableFuncion ( ...... )--  ]

	TEMPORAL_SALTO := analizador.GeneradorGlobal.ObtenerTemporal()

	SALIDA_FINAL.Codigo += fmt.Sprintf("%s = SP + %d; /* SALTANDO A SIGUIENTE ENTORNO*/\n", TEMPORAL_SALTO, ent.Tamanio)

	seEjecuto, codigoParametros := clonFunc.EjecutarParametros(&Entorno_fun, this.ListaExpresiones, TEMPORAL_SALTO, ent)
	if !seEjecuto {
		return entorno.Result3D{Tipo: entorno.NULL}
	}

	SALIDA_FINAL.Codigo += codigoParametros
	SALIDA_FINAL.Codigo += fmt.Sprintf("/* CAMBIANDO DE ENTORNO */; \n")

	SALIDA_FINAL.Codigo += fmt.Sprintf("SP = SP + %d; /* SUMAR TAMAÑO DE ENTORNO ACTUAL PARA SALTAR A FUNCION*/\n", ent.Tamanio)
	SALIDA_FINAL.Codigo += fmt.Sprintf("%s();\n", this.IdFuncion)
	SALIDA_FINAL.Codigo += fmt.Sprintf("SP = SP - %d; /* RESTAR TAMAÑO DE ENTORNO ACTUAL PARA REGRESAR */\n", ent.Tamanio)

	return SALIDA_FINAL
}

func (this Llamada) Get3D(ent *entorno.Entorno) string {

	VALOR := this.Obtener3D(ent)

	return VALOR.Codigo
}

func (l Llamada) ValidarFuncionGenerada(ent *entorno.Entorno, funcion *Simbolos.Funcion) {

	if funcion.Generado == false {
		funcion.Generado = true
		ent.CambiarFuncionGenerada(funcion.Identificador, *funcion)

		CODIGO := funcion.Get3D(ent)
		analizador.GeneradorGlobal.GenerarFuncion(CODIGO)
	}

}

func (this Llamada) Obtener3DRef(ent *entorno.Entorno) entorno.Result3D {
	return entorno.Result3D{}
}
