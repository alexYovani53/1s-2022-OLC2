package Simbolos

import (
	"OLC2/analizador/ast/definicion"
	"OLC2/analizador/ast/interfaces"
	"OLC2/analizador/entorno"
	"fmt"
	"github.com/colegno/arraylist"
	"reflect"
)

var retornoVal = [6][56]entorno.TipoDato{
	{entorno.INTEGER, entorno.FLOAT, entorno.NULL, entorno.NULL, entorno.NULL, entorno.NULL},
	{entorno.FLOAT, entorno.FLOAT, entorno.NULL, entorno.NULL, entorno.NULL, entorno.NULL},
	{entorno.NULL, entorno.NULL, entorno.STRING, entorno.NULL, entorno.NULL, entorno.NULL},
	{entorno.NULL, entorno.NULL, entorno.NULL, entorno.BOOLEAN, entorno.NULL, entorno.NULL},
	{entorno.NULL, entorno.NULL, entorno.NULL, entorno.NULL, entorno.NULL, entorno.NULL},
	{entorno.NULL, entorno.NULL, entorno.NULL, entorno.NULL, entorno.NULL, entorno.VOID},
}

type Funcion struct {
	entorno.Simbolo
	ListaParamsDecl    *arraylist.List
	ListaInstrucciones *arraylist.List
}

func NewFuncion(nombre string, listaParams *arraylist.List, listaInstrucciones *arraylist.List, tipo entorno.TipoDato) Funcion {
	funcSimbolo := entorno.NewSimboloFuncion(0, 0, nombre, tipo, listaParams)

	return Funcion{
		ListaInstrucciones: listaInstrucciones,
		ListaParamsDecl:    listaParams,
		Simbolo:            funcSimbolo,
	}
}

func (f Funcion) EjecutarParametros(ent entorno.Entorno, expresiones *arraylist.List, entRef *entorno.Entorno) bool {

	declaraciones := f.ListaParamsDecl.Clone()

	if declaraciones.Len() != expresiones.Len() {
		fmt.Println("Error en variables")
		return false
	}

	for i := 0; i < declaraciones.Len(); i++ {

		pivoteDec := declaraciones.GetValue(i).(*definicion.Declaracion)
		if pivoteDec.Referencia {
			pivoteDec.EntornoRef = entRef
		}
		pivoteDec.ValorInicializacion = expresiones.GetValue(i).(interfaces.Expresion)

		pivoteDec.Ejecutar(ent)
	}

	return true
}

func (f Funcion) Ejecutar(ent entorno.Entorno) interface{} {

	for i := 0; i < f.ListaInstrucciones.Len(); i++ {

		instrPivote := f.ListaInstrucciones.GetValue(i).(interfaces.Instruccion)

		valorInstruccion := instrPivote.Ejecutar(ent)

		tipoRetFuncion := f.Tipo

		/* El retoro de la función es diferente de VOID y el valor de instruccion es diferente de nil*/

		if valorInstruccion != nil {

			if reflect.TypeOf(valorInstruccion) != reflect.TypeOf(entorno.ValorType{}) {
				fmt.Println("Error en función, se esperaba un retorno valido")
				return entorno.ValorType{Tipo: entorno.NULL, Valor: -1}
			}

			valorRetorno := valorInstruccion.(entorno.ValorType)

			compararTipos := retornoVal[tipoRetFuncion][valorRetorno.Tipo]

			if compararTipos == entorno.NULL {
				fmt.Println("Error tipos, se esperaba un retorno de igual tipo")
				return entorno.ValorType{Tipo: entorno.NULL, Valor: -1}
			}

			return valorRetorno
		}

	}
	return entorno.ValorType{Tipo: entorno.NULL, Valor: -1}
}

/*
fn is_even(a:u8) -> bool {
    if a % 2 == 0 {
        return true
    }
    return false // el valor de retorno en todos los otros casos
}

*/
